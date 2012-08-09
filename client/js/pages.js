/**
 * @fileoverview Utilitaire permettant de faire le passage dynamic de page.
 */
goog.provide('studash.Dashboard')
 
goog.require('goog.net.XhrIo');
goog.require('goog.json');
goog.require('goog.events');
goog.require('goog.dom');
goog.require('goog.dom.classes');

goog.require('studash.fx');

 
/**
 * Fonction construisant le dashboard.
 */
studash.Dashboard.Create =  function() {
	studash.Dashboard.correctMenu(goog.dom.getElementsByClass('navigation'));
	
	// Modify central page content
	var central = goog.dom.createDom('div', {'class' : 'centralContent'});
	studash.Dashboard.createUserArea(central);
	studash.Dashboard.createActions(central);
	studash.Dashboard.createNews(central);
	studash.Dashboard.createAds(central);
	studash.Dashboard.createArrows(central);
	goog.dom.replaceNode(central, goog.dom.getElementByClass('centralContent'));
	
	// Add footer and calendar
  studash.Dashboard.createFooter(document.body);
	studash.Dashboard.createCalendar(document.body);
	
	// Register listener
	goog.events.listen(document.getElementById('footerCalendar'), goog.events.EventType.CLICK, function(e) {
		studash.fx.Fade(
		  goog.dom.getElementByClass('centralContent'),
			studash.fx.FADE_IN, studash.fx.FADE_OUT, .6
	  );
		studash.fx.Fade(
		  goog.dom.getLastElementChild(document.body),
			studash.fx.NEARLY_FADE_OUT, studash.fx.FADE_IN, 1
	  );
	});
};

/**
 * Create User Area
 */
studash.Dashboard.createUserArea =  function(central) {
	central.appendChild(
	  goog.dom.createDom('div', {'class' : 'stdbox'},
      studash.Dashboard.createAvatar(),
      studash.Dashboard.createUserDesc()
    )
  );
};

/**
 * Correct Menu
 * @param {array} nodes du menu
 */
studash.Dashboard.correctMenu =  function(menuNodes) {
	goog.dom.classes.add(menuNodes[0], 'transcript');
	goog.dom.classes.add(menuNodes[1], 'questions');
	goog.dom.classes.add(menuNodes[3], 'clubs');
	goog.dom.classes.add(menuNodes[4], 'moodle');
};

/**
 * Create avatar section
 */
studash.Dashboard.createAvatar =  function() {	
	var avatar = goog.dom.createElement('div');
	var img = goog.dom.createElement('img');
	var span = goog.dom.createElement('span');
	
	avatar.className = 'avatar';
	img.src = 'img/avatar.png';
	span.innerHTML = '1437974';
	
	avatar.appendChild(img);
	avatar.appendChild(goog.dom.createElement('br'));
	avatar.appendChild(span);
	return avatar;
};

/**
 * Create user description section
 */
studash.Dashboard.createUserDesc =  function() {	
	var user = goog.dom.createElement('div');
	goog.dom.classes.add(user, 'userDesc');
	
	// Name
	var name = goog.dom.createElement('span');
	goog.dom.classes.add(name, 'name');
	name.appendChild(goog.dom.createTextNode('Maxime Lavigne'));
	user.appendChild(name);
	user.appendChild(goog.dom.createElement('br'));
	
	// Desc
	var desc = goog.dom.createElement('span');
	var genie = goog.dom.createElement('button');
	var parcour = goog.dom.createElement('button');
	genie.disabled = true;
	parcour.disabled = true;
	genie.appendChild(goog.dom.createTextNode('Génie :'));
	parcour.appendChild(goog.dom.createTextNode('Parcours :'));
	desc.appendChild(genie);
	desc.appendChild(goog.dom.createTextNode('Logiciel'));
	desc.appendChild(parcour);
	desc.appendChild(goog.dom.createTextNode('Régulier'));
	user.appendChild(desc);
	
	return user;
};

/**
 * Create actions section
 */
studash.Dashboard.createActions =  function(central) {
  var actions = ["Nouvelles","Renseignement Personnels","Finances","Horaires","Autres"];	
	var list = goog.dom.createDom('ul', {'class' : 'quicklinks'});
	
	for (idx in actions) {
	  var li = goog.dom.createDom('li', {'class' : 'stdbox'},
		  goog.dom.createDom('h4', null, goog.dom.createTextNode(actions[idx])));
		list.appendChild(li);
	}
	
	central.appendChild(list);
};

/**
 * Create news section
 */
studash.Dashboard.createNews =  function(central) {
  central.appendChild(
	  goog.dom.createDom('div', {'class' : 'stdbox news'},
		  goog.dom.createDom('div', null,
		    goog.dom.createDom('h1', null,
		      goog.dom.createTextNode('Nouvelles!')
		    )
		  )
		)
	);
};

/**
 * Create ads section
 */
studash.Dashboard.createAds =  function(central) {
  central.appendChild(
	  goog.dom.createDom('div', {'class' : 'stdbox ads'},
		  goog.dom.createDom('div', null,
		    goog.dom.createDom('h1', null,
		      goog.dom.createTextNode('Publicités de comités!')
		    )
		  )
		)
	);
};

/**
 * Create footer
 */
studash.Dashboard.createFooter =  function(central) {
  central.appendChild(
	  goog.dom.createDom('footer', null,
		  goog.dom.createDom('div', {'id' : 'footerTask'}, goog.dom.createTextNode('Task')),
			goog.dom.createDom('div', {'id' : 'footerCalendar'}, goog.dom.createTextNode('Calendar')),
			goog.dom.createDom('div', {'id' : 'footerHangout'}, goog.dom.createTextNode('Hangout'))
		)
	);
};

/**
 * Create calendar
 */
studash.Dashboard.createCalendar =  function(central) {			
  var data = {
	  'style' : 'border-width:0',
		'frameborder' : '0',
		'scrolling' : 'no',
		'src' : 'https://www.google.com/calendar/embed?showTitle=0&showTz=0&height=768&wkst=1&bgcolor=%23FFFFFF&src=duguigne%40gmail.com&color=%230D7813&ctz=America%2FMontreal'
	}
				
  central.appendChild(
	  goog.dom.createDom('iframe',	data)
	);
};

/**
 * Create arrows
 */
studash.Dashboard.createArrows =  function(central) {
  central.appendChild(goog.dom.createDom('button', {'class' : 'navArrow', 'id' : 'leftNav'}, goog.dom.createTextNode('<')));
	central.appendChild(goog.dom.createDom('button', {'class' : 'navArrow', 'id' : 'rightNav'}, goog.dom.createTextNode('>')));
};