/* File containing basic functions and reprensentation
 * for the use of studash.
 * Author : Maxime Lavigne
 * Author : Eliott Mahout
 */
 
 var Student = {
   name        : "John Doe",
	 username    : "malesd",
	 password    : "rr11ee22",
	 dateOfBirth : "890822",
	 Credentials : function() { return { code : this.username, nip : this.password, naissance : this.dateOfBirth }},
	 Uri         : function() { return "/u/" + this.username; },
	 TryAuth     : function() {
	   $.post("/auth", this.Credentials()).success(function(data) { console.log(data); });
	 }
 }
 
function FetchActions() {
	$.getJSON(Student.Uri(), function(actions) {
	  
		var nActions = actions.List.length;
		for (i = 0; i < nActions; i++) {
	    var fetchUri = Student.Uri() + "/" + actions.List[i];
		  $('#Informations').append("<div id=\"" + actions.List[i] + "\">Loading...</div>"); 
			$.get(fetchUri)
       .success(function(data) {
			   var action = this.url.replace(Student.Uri() + '/', '');
			   $('#' + action).text(data);
			 });
		}
  });
}