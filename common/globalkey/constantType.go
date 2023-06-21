package globalkey

var DramaTypeRadio int64 = 1
var DramaTypeSound int64 = 2

var DramaStatusNormal int64 = 1
var DramaStatusDown int64 = 2
var DramaStatusShield int64 = 3

var CataLog = map[int]string{
	1: "广播剧",
	2: "有声书",
}

var Category = map[int]string{
	1:  "原创",
	2:  "言情",
	3:  "原创-古代",
	4:  "原创-现代",
	5:  "原创-悬疑",
	6:  "原创-幻想",
	7:  "原创-架空",
	8:  "原创-百合",
	9:  "言情-古代",
	10: "言情-现代",
	11: "言情-悬疑",
	12: "言情-幻想",
}

var Process = map[int]string{
	1: "连载",
	2: "完结",
}

var Status = map[int]string{
	1: "正常",
	2: "下架",
	3: "屏蔽",
}

var TempType = map[int]string{
	1: "单排",
	2: "横排",
	3: "竖排",
}

var ProducerStatus = map[int]string{
	1: "正常",
	2: "停用",
}
