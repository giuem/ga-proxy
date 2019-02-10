(function(win, doc, navigator) {
  var screen = win.screen;
  var encode = encodeURIComponent;
  var max = Math.max;
  // const min = Math.min;
  var performance = win.performance;
  var timing = performance && performance.timing;

  var pvData = {
    ga: win.ga_tid,
    dt: doc.title,
    de: doc.characterSet || doc.charset,
    dr: doc.referrer || void 0,
    ul:
      navigator.language ||
      navigator.browserLanguage ||
      navigator.userLanguage ||
      void 0,
    sd: screen.colorDepth + "-bit",
    sr: screen.width + "x" + screen.height,
    vp:
      max(doc.documentElement.clientWidth, win.innerWidth || 0) +
      "x" +
      max(doc.documentElement.clientHeight, win.innerHeight || 0),
    z: Date.now()
  };

  function buildQueryString(params) {
    var qs = [];
    for (var k in params) {
      if (params.hasOwnProperty(k) && params[k] !== void 0) {
        qs.push(encode(k) + "=" + encode(params[k]));
      }
    }
    return qs.join("&");
  }

  function sendViaImg(uri, params) {
    var img = new Image();
    // img.width = img.height = 1;
    img.src = win.ga_url + uri + "?" + buildQueryString(params);
  }

  function sendTiming() {
    if (!timing) { return; }
    var navigationStart = timing.navigationStart;
    if (navigationStart == 0) { return; }

    var filterNumber = function (num) { return isNaN(num) || num == Infinity || num < 0 ? void 0 : num; };

    var perfData = {
      plt: filterNumber(timing.loadEventStart - navigationStart),
      dns: filterNumber(timing.domainLookupEnd - timing.domainLookupStart),
      pdt: filterNumber(timing.responseEnd - timing.responseStart),
      rrt: filterNumber(timing.redirectEnd - timing.redirectStart),
      tcp: filterNumber(timing.connectEnd - timing.connectStart),
      srt: filterNumber(timing.responseStart - timing.requestStart),
      dit: filterNumber(timing.domInteractive - navigationStart),
      clt: filterNumber(timing.domContentLoadedEventStart - navigationStart)
    };

    for (var key in pvData) {
      perfData[key] = pvData[key];
    }

    sendViaImg("/t", perfData);
  }

  // page view
  sendViaImg("/p", pvData);
  // timing
  if (document.readyState == "complete") {
    sendTiming();
  } else {
    win.addEventListener("load", sendTiming);
  }
})(window, document, navigator);
