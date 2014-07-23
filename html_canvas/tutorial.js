var draw_tutorial = function() {
    var tutorial_canvas = document.getElementById("tutorial");
    var tutorial_context = tutorial_canvas.getContext("2d");
	alert("draw method");
    tutorial_canvas.fillRect(50,25,150,100);
}

$(document).ready(function() {
  //var canvas = $("#tutorial");
  var canvas = document.getElementById("tutorial");
  var ctx = canvas.getContext('2d');

  if(canvas.getContext){
	  var ctx = canvas.getContext('2d');
	  //drawing
  } else {
	  //canvas-unsupported
  }

});
