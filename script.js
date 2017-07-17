  var ga = "UA-102578664-1";
(function(ga, url,window, document, navigator, location) {
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
    'z=' + (+new Date)
  ];

  window.__beacon_img = new Image();
  window.__beacon_img.src = url + '?' + data.join('&');;
})(ga, "https://ga.giuem.com",window, document, navigator, location);
