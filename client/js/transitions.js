/**
 * @fileoverview Utilitaire permettant de faire des transitions.
 */
goog.provide('studash.fx')
 
goog.require('goog.events');
goog.require('goog.dom');
goog.require('goog.dom.classes');
goog.require('goog.fx.css3.Transition');
 
studash.fx.FADE_OUT = { 'opacity' : '0' }
studash.fx.NEARLY_FADE_OUT = { 'opacity' : '0.3' }
studash.fx.FADE_IN = { 'opacity' : '1.0' }

studash.fx.Fade = function(elem, start, end, timeInSec) {
	(new goog.fx.css3.Transition(elem, 1, start, end, [
    {'property': 'opacity', 'duration': timeInSec, 'timing': 'ease-in', 'delay': 0}
  ])).play();
}
