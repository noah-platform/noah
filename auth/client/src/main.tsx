import React from "react";
import { createRoot } from "react-dom/client";
import App from "./app";

import "./globals.css";
import "@noah/ui/styles.css";

const el = document.getElementById("root");
if (el) {
  const root = createRoot(el);
  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>
  );
} else {
  throw new Error("Could not find root element");
}
