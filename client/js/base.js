/*
 * Basic student reprensentation
 */
 
 var Student = {
   name        : "John Doe",
	 username    : "malesd",
	 password    : "rr11ee22",
	 dateOfBirth : "890822",
	 Uri         : function() { return "/u/" + this.username; }
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