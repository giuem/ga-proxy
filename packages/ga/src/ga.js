(function(win, doc, navigator) {
  const screen = win.screen;
  const encode = encodeURIComponent;
  const max = Math.max;
  // const min = Math.min;
  const performance = win.performance;
  const timing = performance && performance.timing;

  const pvData = {
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
    const qs = [];
    for (const k in params) {
      if (params.hasOwnProperty(k) && params[k] !== void 0) {
        qs.push(encode(k) + "=" + encode(params[k]));
      }
    }
    return qs.join("&");
  }

  function sendViaImg(uri, params) {
    const img = new Image();
    // img.width = img.height = 1;
    img.src = win.ga_url + uri + "?" + buildQueryString(params);
  }

  function sendTiming() {
    if (!timing) return;
    const navigationStart = timing.navigationStart;
    if (navigationStart == 0) return;

    const filterNumber = num =>
      isNaN(num) || num == Infinity || num < 0 ? void 0 : num;

    const perfData = {
      plt: filterNumber(timing.loadEventStart - navigationStart),
      dns: filterNumber(timing.domainLookupEnd - timing.domainLookupStart),
      pdt: filterNumber(timing.responseEnd - timing.responseStart),
      rrt: filterNumber(timing.redirectEnd - timing.redirectStart),
      tcp: filterNumber(timing.connectEnd - timing.connectStart),
      srt: filterNumber(timing.responseStart - timing.requestStart),
      dit: filterNumber(timing.domInteractive - navigationStart),
      clt: filterNumber(timing.domContentLoadedEventStart - navigationStart)
    };

    for (const key in pvData) {
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
