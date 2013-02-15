jQuery.ajaxSetup({cache:false});
$("document").ready(function(){
retrieveContactList();
	$("#frmContacts").change(function(){
		retriveContact();
	});
});
var saveContact = function(){
	clearNotices();
	var strId = $("#frmContacts").find(":selected").val();
	var strName = addslashes($("#frmName").val());
	if((strName == "")||(strName.search(/\S/) == -1)){
		notice("A name is required to save a new contact.");
		return;
	}
	var strAddress = addslashes($("#frmAddress").val());
	var strPhone = addslashes($("#frmPhone").val());
	var strEmail = addslashes($("#frmEmail").val());
	
	console.lo
	
	var url = "/request/saveContact";
	console.log("save");
	$.post(url, 
		{ 'id': strId, 'name': strName, 'address': strAddress,
		'phone': strPhone, 'email': strEmail },
		function(data){
			retrieveContactList();
			if(data == 'same'){
				notice("A duplicate entry already exists");
			}
		});
}

var deleteContact = function(){
	clearNotices();
	var strId = $("#frmContacts").find(":selected").val();
	if(strId == '-1'){
		notice("This is not an entry and cannot be deleted");
		return;
	}
	var url = "/request/deleteContact";
	$.post(url, { 'id': strId }, function(data){
		retrieveContactList();
	});
}

var retrieveContactList = function(){
	var url = "/request/retrieveList";
	$.get(url, function(data, textStatus, jqXHR){
// 		console.log(data)
		$("#frmContacts").html(data);
		retriveContact();
	});
}

var retriveContact = function(){
	clearNotices();
	var strId = $("#frmContacts").find(":selected").val();
	// console.log("ID: "+strId);
	var url = "/request/retrieveContact";
	$.post(url, { 'id': strId }, function(data){
// 		console.log(data);
		if(data != -1){
				myData = JSON.parse(data);
// 				console.log(myData)
				$("#frmName").val(myData['fldName']);
				$("#frmAddress").val(myData['fldAddress']);
				$("#frmPhone").val(myData['fldPhone']);
				$("#frmEmail").val(myData['fldEmail']);
			}else{
				$("#frmName").val("");
				$("#frmAddress").val("");
				$("#frmPhone").val("");
				$("#frmEmail").val("");
			}
		});
}

var notice = function(msg){
	$("#alerts").append("<li>"+msg+"</li>");
}

var clearNotices = function(){
	$("#alerts").html("");
}

function addslashes(string) {
    return string.replace(/\\/g, '\\\\').
        replace(/\u0008/g, '\\b').
        replace(/\t/g, '\\t').
        replace(/\n/g, '\\n').
        replace(/\f/g, '\\f').
        replace(/\r/g, '\\r').
        replace(/'/g, '\\\'').
        replace(/"/g, '\\"');
}