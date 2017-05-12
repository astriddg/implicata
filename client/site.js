(function() {
	var elapsed = 0
	var resizeFromWidth = window.innerWidth
	var resizeFromHeight = window.innerHeight
	var startTime
	var resized = false

	var constants = {
		website_url: window.location.href,
		session_id: guid(),
	}

	function guid() {
		function s4() {
			return Math.floor((1 + Math.random()) * 0x10000).toString(16).substring(1)
		}
		return s4() + s4() + '-' + s4() + '-' + s4() + '-' + s4() + '-' + s4() + s4() + s4()
	}
	
	function copyPaste(event) {
		var fields = {}
		fields[event.target.id] = true
		var data = {
			copy_and_paste: fields
		}	
		post(data)
	}

	function startTimer() {
		if (!startTime) {
			startTime = Date.now()
		}
	}	

	function stopTimer() {
		if (!startTime) return 0
		var elapsed = (Date.now() - startTime) / 1000
		startTime = undefined
		return Math.round(elapsed)
	}

	function resize(event) {
		if (resized) return
		resized = true

		var data = {
			resize_from: {
				height: String(resizeFromHeight),
				width: String(resizeFromWidth)
			},
			resize_to: {
				width: String(event.target.innerWidth),
				height: String(event.target.innerHeight)
			}
		}
		post(data)
	}

	function submit() {
		var elapsed = stopTimer()
		var data = {
			form_completion_time: elapsed
		}
		post(data)
	}

	function post(data) {
		var body = Object.assign({}, data, constants)
		fetch('http://localhost:8080/submit', {
			method: 'POST',
			body: JSON.stringify(body)
		}).then(response => {
			if (response && response.status !== 204) {
				alert(response.statusText)
			}
		})
	}

	var emailAddress = document.getElementById('inputEmail')
	emailAddress.addEventListener('keyup', startTimer)
	emailAddress.addEventListener('paste', copyPaste)

	var cardNumber = document.getElementById('inputCardNumber')
	cardNumber.addEventListener('keyup', startTimer)
	cardNumber.addEventListener('paste', copyPaste)

	var securityCode = document.getElementById('inputCVV')
	securityCode.addEventListener('keyup', startTimer)
	securityCode.addEventListener('paste', copyPaste)

	var form = document.getElementById('form')
	form.addEventListener('submit', submit)

	window.addEventListener('resize', resize, true)
})()
