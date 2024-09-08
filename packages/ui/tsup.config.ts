import { defineConfig } from "tsup";

export default defineConfig((options) => ({
  entryPoints: ["src/**/*"],
  format: ["esm"],
  dts: true,
  sourcemap: true,
  bundle: false,
  external: ["react"],
  ...options,
}));
