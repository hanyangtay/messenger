function SlidingHeader(options) {
	this.init(options);
}

SlidingHeader.prototype.init = function(options) {
	// Define the settings as an object
	var settings = {
		element: options.element,
		class: 'to_scroll'
	}

	// If the target element has a specific class declared in the object passed,
	// replace the default class key in settings object
	if (options.class) {
		settings.class = options.class;
	}

	// Get the target element from the DOM
	var elementToSlide = document.querySelector(settings.element);

	// If there's a valid element, slide it. If not, just return false
	if (elementToSlide != null) {
		var elementSettings = {
			elementHeight: elementToSlide.clientHeight,
			class: settings.class
		}

		this.slide(elementToSlide, elementSettings);
	} else {
		return false;
	}
}

SlidingHeader.prototype.slide = function(element, elementSettings) {
	
	var offset = 0,
		lastPosition = 0,
		targetClass = elementSettings.class,
		elementHeight = elementSettings.elementHeight;

	window.addEventListener('scroll', function(e) {
		
		var newPosition = this.scrollY,
			position = newPosition - lastPosition;			

		// If we scrolled more than the element's height, then add the class to it.
		// Else, remove it, and the element will show up again
		if (offset + position > elementHeight) {
			offset = elementHeight;

			if (!element.classList.contains(targetClass)) {
				element.className = 'navigation-container scrolled';
			}
		} else {
			offset = offset + position;

			if (element.classList.contains(targetClass)) {
				element.className = 'navigation-container';
			}
		}

		if (offset < 0) {
			offset = 0;
		}

		lastPosition = newPosition;

	});
}

export default SlidingHeader




