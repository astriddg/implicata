'use strict';

var opts,
Site = {
	options: {
		emailAddress: document.getElementById('inputEmail'),
		cardNumber: document.getElementById('inputCardNumber'),
		securityCode: document.getElementById('inputCVV'),
	},

	init: function() {
		opts = this.options;
	}
};

(function() {
	Site.init();
})();
