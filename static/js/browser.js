// Check for Ogg

var temp = document.createElement('audio');

isSupp = temp.canPlayType('audio/ogg');

if (isSupp == "") {
	var supportOgg = false;
} else {
	var supportOgg = true;
}

// Check for Webm

var temp = document.createElement('video');

isSupp = temp.canPlayType('video/webm');

if (isSupp == "") {
	var supportWebm = false;
} else {
	var supportWebm = true;
}

// Check for Opus

var temp = document.createElement('audio');

isSupp = temp.canPlayType('audio/opus');

if (isSupp == "") {
	var supportOpus = false;
} else {
	var supportOpus = true;
}

// 1.1: Check for flac

var temp = document.createElement('audio');
isSupp = temp.canPlayType('audio/x-flac');
if (isSupp == "") {
	var supportFlac = false;
} else {
	var supportFlac = true;
}

// Check browser

var userAgent = navigator.userAgent;
var browser = "none";
var browserChrome = false;

if (userAgent.indexOf('Firefox') > -1) {
	browser = "Firefox";

} else if (userAgent.indexOf('Trident') > -1) {
	browser = "rv";

} else if (userAgent.indexOf('Safari') > -1) {

	if (userAgent.indexOf('Edge') > -1) {
		browser =  "Edge";
		browserChrome = true;
	} else if (userAgent.indexOf('OPR') > -1) {
		browser = "OPR";
		browserChrome = true;
	} else if (userAgent.indexOf('Vivaldi') > -1) {
		browser = "Vivaldi";
		browserChrome = true;
	} else if (userAgent.indexOf('Chrome') > -1) {
		browser = "Chrome";
	} else {
		browser = "Version";
	}

} else if (userAgent.indexOf('Silk') > -1) {
	browser = "Silk";

} else {
	browser = "Unknown";
}

// Check browser version

var browserVer = 0;

browserVerPre = userAgent.substring(userAgent.indexOf(browser) + (browser.length + 1));
browserVer = browserVerPre.substring(0, browserVerPre.indexOf("."));

browserChromeVerPre = userAgent.substring(userAgent.indexOf('Chrome') + 7);
browserChromeVer = browserChromeVerPre.substring(0, browserChromeVerPre.indexOf("."));

// Check for IE, Safari or Opera and convert

if (browser == "Version") { browser = "Safari"; }
if (browser == "rv") { browser = "Internet Explorer"; }
if (browser == "OPR") { browser = "Opera"; }