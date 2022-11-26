import vue from "@vitejs/plugin-vue";
import { defineConfig } from "vite";
import { quasar, transformAssetUrls } from "@quasar/vite-plugin";
import { fileURLToPath, URL } from "url";

export default defineConfig({
  plugins: [
    vue({
      template: { transformAssetUrls },
    }),
    quasar({
      sassVariables: "src/quasar-variables.sass",
    }),
  ],
  server: {
    port: 3000,
  },
  // preview: {
  //   port: 3000,
  // },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
