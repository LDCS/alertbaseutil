package alertbaseutil


import (
	"fmt"
	"strings"
	"time"
	"github.com/LDCS/genutil"
)

func (self *ROW) GetMap() map[string]string {
	return map[string]string{
		"subtab"               :    self.Subtab,
		"level"                :    self.Level,
		"subject"              :    self.Subject,
		"escalate"             :    self.Escalate,
		"escalate-minutes1"    :    self.Escalatemin1,
		"escalate-minutes2"    :    self.Escalatemin2,
		"subjectnum"           :    self.Subjectnum,
		"doneat"               :    self.Doneat,
		"openat"               :    self.Openat,
		"owner"                :    self.Owner,
		"assigner"             :    self.Assigner,
		"status"               :    self.Comment,
		"comment"              :    self.Comment,
	}
}

func (self *ROW) SetFromKVL(kvplist string) {
	kvplist = strings.Replace(kvplist, ",", ".", -1)
	mp := genutil.GetMapFromKV(kvplist)
	self.Subtab                =          mp["subtab"]
	self.Level                 =          mp["level"]
	self.Subject               =          mp["subject"]
	self.Escalate              =          mp["escalate"]
	self.Escalatemin1          =          mp["escalate-minutes1"]
	self.Escalatemin2          =          mp["escalate-minutes2"]
	self.Subjectnum            =          mp["subjectnum"]
	self.Doneat                =          mp["doneat"]
	self.Openat                =          mp["openat"]
	self.Owner                 =          mp["owner"]
	self.Assigner              =          mp["assigner"]
	self.Status                =          mp["status"]
	self.Comment               =          mp["comment"]
}

func (self *ROW) GetCSV() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		self.Subtab,
		self.Level,
		self.Subject,
		self.Escalate,
		self.Escalatemin1,
		self.Escalatemin2,
		self.Subjectnum,
		self.Doneat,
		self.Openat,
		self.Owner,
		self.Assigner,
		self.Status,
		self.Comment,
	)
}

func (self *OPSROW) GetCSV() string {
	return fmt.Sprintf("%s,%s", self.Opsname, self.Takenoverat)
}

func (self *ROW) GetHeader() string {
	return "subtab,level,subject,escalate,escalate-minutes1,escalate-minutes2,subjectnum,doneat,openat,owner,assigner,status,comment"
}

func (self *OPSROW) GetHeader() string {
	return "username,takenoverat"
}

func (self *ROW) SetFromCSV(csvline string) error {
	fields := strings.Split(csvline, ",")
	if len(fields) != 13 { return fmt.Errorf("Bad csv line") }
	
	self.Subtab,
	self.Level,
	self.Subject,
	self.Escalate,
	self.Escalatemin1,
	self.Escalatemin2,
	self.Subjectnum,
	self.Doneat,
	self.Openat,
	self.Owner,
	self.Assigner,
	self.Status,
	self.Comment = fields[0], fields[1], fields[2], fields[3], fields[4], fields[5], fields[6], fields[7], fields[8], fields[9], fields[10], fields[11], fields[12]
	return nil
}

func (self *OPSROW) SetFromCSV(csvline string) error {
	fields := strings.Split(csvline, ",")
	if len(fields) != 2 { return fmt.Errorf("Bad csv line") }
	self.Opsname, self.Takenoverat = fields[0], fields[1]
	return nil
}

func (self *ROW) GetKey() int64 {
	return genutil.ToInt(self.Openat, int64(-1))
}

func (self *ROW) GetOpenat() time.Time {
	return time.Unix(0, self.GetKey()*1000000)
}

func (self *ROW) IsOlderThan(d time.Duration) bool {
	tstamp := time.Unix(0, self.GetKey()*1000000)
	return (time.Now().Sub(tstamp) > d)
}

func (self *ROW) UpdateWith(targrow *ROW) {
	if self.Openat != targrow.Openat { return }
	if targrow.Subtab       != "" { self.Subtab       =      targrow.Subtab       }
	if targrow.Level        != "" { self.Level        =      targrow.Level        }
	if targrow.Subject      != "" { self.Subject      =      targrow.Subject      }
	if targrow.Escalate     != "" { self.Escalate     =      targrow.Escalate     }
	if targrow.Escalatemin1 != "" { self.Escalatemin1 =      targrow.Escalatemin1 }
	if targrow.Escalatemin2 != "" { self.Escalatemin2 =      targrow.Escalatemin2 }
	if targrow.Subjectnum   != "" { self.Subjectnum   =      targrow.Subjectnum   }
	if targrow.Doneat       != "" { self.Doneat       =      targrow.Doneat       }
	if targrow.Owner        != "" { self.Owner        =      targrow.Owner        }
	if targrow.Assigner     != "" { self.Assigner     =      targrow.Assigner     }
	if targrow.Status       != "" { self.Status       =      targrow.Status       }
	if targrow.Comment      != "" { self.Comment      =      targrow.Comment      }
	
}



func (self ROWS) Len() int {
	return len(self)
}

func (self ROWS) Less(ii, jj int) bool {
	return self[ii].GetKey() > self[jj].GetKey()
}

func (self ROWS) Swap(ii ,jj int) {
	var tmp *ROW
	tmp = self[ii]
	self[ii] = self[jj]
	self[jj] = tmp
}
