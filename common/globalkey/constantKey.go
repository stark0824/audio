package globalkey

/**
global constant key
*/

// DelStateNo 软删除
var DelStateNo int64 = 0  //未删除
var DelStateYes int64 = 1 //已删除

var StateNormal int64 = 1
var StateDelete int64 = 0

// DateTimeFormatTplStandardDateTime 时间格式化模版
var DateTimeFormatTplStandardDateTime = "Y-m-d H:i:s"
var DateTimeFormatTplStandardDate = "Y-m-d"
var DateTimeFormatTplStandardTime = "H:i:s"
