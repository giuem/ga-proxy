var ga = "UA-xxxx-x";
var ga_url = "https://ga.giuem.com";
(function(ga, url, window, document, navigator, location) {
  var screen = window.screen;
  var encode = encodeURIComponent;

  var data = [
    'ga=' + ga,
    'dt=' + encode(document.title),
    'dr=' + encode(document.referrer),
    'ul=' + (navigator.language || navigator.browserLanguage || navigator.userLanguage),
    'sd=' + screen.colorDepth + '-bit',
    'sr=' + screen.width + 'x' + screen.height,
    'vp=' + Math.max(document.documentElement.clientWidth, window.innerWidth || 0) + 'x' + Math.max(document.documentElement.clientHeight, window.innerHeight || 0),
    'z=' + Date.now()
  ];

  window.__ga_img = new Image();
  window.__ga_img.src = url + '?' + data.join('&');;
})(ga, ga_url, window, document, navigator, location);
