window.onload = function() {
  window.ui = SwaggerUIBundle({
    urls: ###JSONLIST###,
    dom_id: '#swagger-ui',  
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });
};
