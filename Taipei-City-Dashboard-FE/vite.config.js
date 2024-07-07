import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import viteCompression from "vite-plugin-compression";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [vue(), viteCompression()],
	build: {
		rollupOptions: {
			output: {
				manualChunks(id) {
					if (id.includes("node_modules")) {
						return id
							.toString()
							.split("node_modules/")[1]
							.split("/")[0]
							.toString();
					}
				},
			},
		},
		chunkSizeWarningLimit: 1600,
	},
	base: "/",
	server: {
		host: "0.0.0.0",
		port: 80,
		proxy: {
			"/api/dev": {
				target: "http://dashboard-be:8080",
				changeOrigin: true,
				rewrite: (path) => path.replace("/dev", "/v1"),
			},
			"/geo_server": {
				target: "https://geoserver.tuic.gov.taipei/geoserver/",
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/geo_server/, ""),
			},
			proxy: {
				"/api/v1": {
					target: "http://localhost:8088",
					changeOrigin: true,
					rewrite: (path) => path.replace(/^\/api\/v1/, ""),
				},
			},
		},
	},
});

// import { defineConfig } from "vite";
// import vue from "@vitejs/plugin-vue";
// import viteCompression from "vite-plugin-compression";
// import { createProxyMiddleware } from "http-proxy-middleware";

// export default defineConfig({
// 	plugins: [vue(), viteCompression()],
// 	build: {
// 		rollupOptions: {
// 			output: {
// 				manualChunks(id) {
// 					if (id.includes("node_modules")) {
// 						return id
// 							.toString()
// 							.split("node_modules/")[1]
// 							.split("/")[0]
// 							.toString();
// 					}
// 				},
// 			},
// 		},
// 		chunkSizeWarningLimit: 1600,
// 	},
// 	base: "/",
// 	server: {
// 		host: "0.0.0.0",
// 		port: 80,
// 		proxy: {
// 			"/api": {
// 				target: "http://localhost:8088",
// 				changeOrigin: true,
// 				rewrite: (path) => path.replace(/^\/api/, ""),
// 				configure: (proxy, options) => {
// 					proxy.on("proxyReq", (proxyReq, req, res) => {
// 						proxyReq.setHeader("Origin", "http://localhost:8080");
// 					});
// 				},
// 			},
// 			"/geo_server": {
// 				target: "https://geoserver.tuic.gov.taipei/geoserver/",
// 				changeOrigin: true,
// 				rewrite: (path) => path.replace(/^\/geo_server/, ""),
// 			},
// 		},
// 	},
// });
