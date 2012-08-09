/**
 * @fileoverview Fonctionnalité lié au menu de bas de page.
 */
goog.provide('studash.footer')
 
/**
 * Configuration des trois bouttons.
 * @enum {string}
 */
studash.footer.TASK = {'id': 'footerTask', 'name': 'Tâches', 'active' : false}
studash.footer.CALENDAR = {'id': 'footerCalendar', 'name': 'Calendrier', 'active' : false}
studash.footer.CONVERSE = {'id': 'footerHangout', 'name': 'Converser', 'active' : false}

studash.footer.BEHIND = -5;
studash.footer.FRONT = 5;

studash.footer.Make = function() {
  var b = [studash.footer.TASK, studash.footer.CALENDAR, studash.footer.CONVERSE]
  var footer = goog.dom.createElement('footer');
	for (i in b) {
	  footer.appendChild(
		  goog.dom.createDom('div', {'id' : b[i].id}, goog.dom.createTextNode(b[i].name))
	  );
	}
	return footer;
}

studash.footer.Register = function() {
  var b = [studash.footer.TASK, studash.footer.CALENDAR, studash.footer.CONVERSE]
	for (i in b) {
		goog.events.listen(
			document.getElementById(b[i].id),
			goog.events.EventType.CLICK,
			studash.footer.onClick
	  );
	}
}

studash.footer.onClick = function(event) {
	switch (event.target.id) {
	  case studash.footer.TASK.id:
		  console.log('click TASK')
			break;
	  case studash.footer.CALENDAR.id:
		  studash.footer.onCalendarClick();
			break;
		case studash.footer.CONVERSE.id:
		  console.log('click CONVERSE')
			break;
	}
}

studash.footer.onCalendarClick = function() {
  var c = studash.footer.CALENDAR;
	var i = studash.fx.FADE_IN;
	var o = studash.fx.FADE_OUT;
	var no = studash.fx.NEARLY_FADE_OUT;
	
	var eCalendar = goog.dom.getLastElementChild(document.body);
	var eContent = goog.dom.getElementByClass('centralContent');
	
	studash.fx.Fade(eContent, c.active?o:i, c.active?i:o, .6);
	studash.fx.Fade(eCalendar, c.active?i:no, c.active?no:i, 1);
	
	goog.Timer.callOnce(studash.footer.focusCalendar, 1000)
}

studash.footer.focusCalendar = function() {
  var c = studash.footer.CALENDAR;
	var eCalendar = goog.dom.getLastElementChild(document.body);
	eCalendar.style['z-index'] = c.active ? studash.footer.BEHIND : studash.footer.FRONT;
  c.active = !c.active;
}
		// studash.fx.Fade(
		  // goog.dom.getElementByClass('centralContent'),
			// studash.fx.FADE_IN, studash.fx.FADE_OUT, .6
	  // );
		// studash.fx.Fade(
		  // goog.dom.getLastElementChild(document.body),
			// studash.fx.NEARLY_FADE_OUT, studash.fx.FADE_IN, 1
	  // );