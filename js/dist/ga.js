(function(window, document, navigator) {
  var url = window.ga_url;
  var tid = window.ga_tid;

  var screen = window.screen;
  var encode = encodeURIComponent;
  var max = Math.max;
  var min = Math.min;

  var data = [
    "ga=" + tid,
    "dt=" + encode(document.title),
    "de=" + encode(document.characterSet || document.charset),
    "dr=" + encode(document.referrer),
    "ul=" +
      (navigator.language ||
        navigator.browserLanguage ||
        navigator.userLanguage),
    "sd=" + screen.colorDepth + "-bit",
    "sr=" + screen.width + "x" + screen.height,
    "vp=" +
      max(document.documentElement.clientWidth, window.innerWidth || 0) +
      "x" +
      max(document.documentElement.clientHeight, window.innerHeight || 0),
    "z=" + Date.now()
  ];

  window.__ga_img = new Image();
  window.__ga_img.src = url + "?" + data.join("&");
})(window, document, navigator);
