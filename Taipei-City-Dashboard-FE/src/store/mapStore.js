// Developed by Taipei Urban Intelligence Center 2023-2024

/* mapStore */
/*
The mapStore controls the map and includes methods to modify it.

!! PLEASE BE SURE TO REFERENCE THE MAPBOX DOCUMENTATION IF ANYTHING IS UNCLEAR !!
https://docs.mapbox.com/mapbox-gl-js/guides/
*/
import { createApp, defineComponent, nextTick, ref } from "vue";
import { defineStore } from "pinia";
import mapboxGl from "mapbox-gl";
import "mapbox-gl/dist/mapbox-gl.css";
import axios from "axios";
import { Threebox } from "threebox-plugin";

// Other Stores
import { useAuthStore } from "./authStore";
import { useDialogStore } from "./dialogStore";

// Vue Components
import MapPopup from "../components/map/MapPopup.vue";

// Utility Functions or Configs
import mapStyle from "../assets/configs/mapbox/mapStyle.js";

import * as turf from "@turf/turf";
import {
	MapObjectConfig,
	TaipeiTown,
	TaipeiVillage,
	TaipeiBuilding,
	TpDistrict,
	TpVillage,
	maplayerCommonPaint,
	maplayerCommonLayout,
} from "../assets/configs/mapbox/mapConfig.js";
import { savedLocations } from "../assets/configs/mapbox/savedLocations.js";
import { calculateGradientSteps } from "../assets/configs/mapbox/arcGradient";
import { voronoi } from "../assets/utilityFunctions/voronoi.js";
import { interpolation } from "../assets/utilityFunctions/interpolation.js";
import { marchingSquare } from "../assets/utilityFunctions/marchingSquare.js";

export const useMapStore = defineStore("map", {
	state: () => ({
		// Array of layer IDs that are in the map
		currentLayers: [],
		// Array of layer IDs that are in the map and currently visible
		currentVisibleLayers: [],
		// Stores all map configs for all layers (to be used to render popups)
		mapConfigs: {},
		// Stores the mapbox map instance
		map: null,
		// Stores popup information
		popup: null,
		// Stores saved locations
		savedLocations: savedLocations,
		// Store currently loading layers,
		loadingLayers: [],
		allPoints: [], // This will store features from various GeoJSON files
		currentLocationMarker: null, // Stores the current location marker instance
	}),
	getters: {},
	actions: {
		/* Initialize Mapbox */
		// 1. Creates the mapbox instance and passes in initial configs
		// 3. Set the map center to a new location
		setMapCenter(centerCoordinates, zoomLevel = 13.9) {
			// 可以設定一個默認的縮放級別，比如 15
			if (this.map && centerCoordinates.length === 2) {
				// 使用 flyTo 方法來更新中心點和縮放級別
				this.map.flyTo({
					center: centerCoordinates,
					zoom: zoomLevel, // 縮放級別可以是任何您需要的值，這裡設為 15
					essential: true, // 這個選項是為了在使用 CSS transform 改變地圖大小時保證動畫正常執行
				});

				this.map.setMaxBounds([
					[121.3870596781498, 24.95733863075891], // Southwest coordinates
					[121.6998231749096, 25.21179993640203], // Northeast coordinates
				]);
			} else {
				console.error("Invalid coordinates or map is not initialized");
			}
		},

		// findNearestPoint(currentLocation) {
		// 	if (this.allPoints.length === 0) {
		// 		console.error("No points available to find the nearest.");
		// 		return null;
		// 	}
		// 	const pointsFeatureCollection = turf.featureCollection(
		// 		this.allPoints
		// 	);
		// 	const nearest = turf.nearestPoint(
		// 		turf.point(currentLocation),
		// 		pointsFeatureCollection
		// 	);
		// 	if (nearest) {
		// 		this.showPopupAtPoint(nearest);
		// 	}
		// 	return nearest;
		// },

		// findNearestPoint(currentLocation) {
		// 	if (this.allPoints.length === 0) {
		// 		console.error("No points available to find the nearest.");
		// 		return null;
		// 	}
		// 	const pointsFeatureCollection = turf.featureCollection(
		// 		this.allPoints.map((feature) => ({
		// 			type: "Feature",
		// 			geometry: {
		// 				type: "circle",
		// 				coordinates: [
		// 					feature.coordinates.longitude,
		// 					feature.coordinates.latitude,
		// 				],
		// 			},
		// 			properties: feature.properties,
		// 		}))
		// 	);

		// 	const nearest = turf.nearestPoint(
		// 		turf.point(currentLocation),
		// 		pointsFeatureCollection
		// 	);
		// 	if (nearest) {
		// 		this.showPopupAtPoint(nearest);
		// 	}
		// 	return nearest;
		// },

		// findNearestPoint(currentLocation) {
		// 	if (
		// 		!currentLocation ||
		// 		currentLocation.length < 2 ||
		// 		isNaN(currentLocation[0]) ||
		// 		isNaN(currentLocation[1])
		// 	) {
		// 		console.error(
		// 			"Invalid or missing coordinates",
		// 			currentLocation
		// 		);
		// 		return null;
		// 	}

		// 	if (this.allPoints.length === 0) {
		// 		console.error("No points available to find the nearest.");
		// 		return null;
		// 	}

		// 	const pointsFeatureCollection = turf.featureCollection(
		// 		this.allPoints.map((feature) => ({
		// 			type: "Feature",
		// 			geometry: {
		// 				type: "Point",
		// 				coordinates: [
		// 					feature.coordinates.longitude,
		// 					feature.coordinates.latitude,
		// 				],
		// 			},
		// 			properties: feature.properties,
		// 		}))
		// 	);

		// 	const nearest = turf.nearestPoint(
		// 		turf.point(currentLocation),
		// 		pointsFeatureCollection
		// 	);
		// 	if (nearest) {
		// 		this.showPopupAtPoint(nearest);
		// 	}
		// 	return nearest;
		// },

		// findNearestPoint(currentLocation) {
		// 	if (
		// 		!currentLocation ||
		// 		currentLocation.length < 2 ||
		// 		isNaN(currentLocation[0]) ||
		// 		isNaN(currentLocation[1])
		// 	) {
		// 		console.error(
		// 			"Invalid or missing coordinates",
		// 			currentLocation
		// 		);
		// 		return null;
		// 	}

		// 	if (this.allPoints.length === 0) {
		// 		console.error("No points available to find the nearest.");
		// 		return null;
		// 	}

		// 	console.log("Current Location:", currentLocation);
		// 	console.log("All Points:", this.allPoints);

		// 	const pointsFeatureCollection = turf.featureCollection(
		// 		this.allPoints.map((feature) => ({
		// 			type: "Feature",
		// 			geometry: {
		// 				type: "Point",
		// 				coordinates: feature.coordinates,
		// 			},
		// 			properties: feature.properties,
		// 		}))
		// 	);

		// 	const nearest = turf.nearestPoint(
		// 		turf.point(currentLocation),
		// 		pointsFeatureCollection
		// 	);
		// 	console.log("Nearest Feature:", nearest);
		// 	if (nearest) {
		// 		this.showPopupAtPoint(nearest);
		// 	}
		// 	return nearest;
		// },

		findNearestPoint(currentLocation) {
			if (
				!currentLocation ||
				currentLocation.length < 2 ||
				isNaN(currentLocation[0]) ||
				isNaN(currentLocation[1])
			) {
				console.error(
					"Invalid or missing coordinates",
					currentLocation
				);
				return null;
			}

			if (this.allPoints.length === 0) {
				console.error("No points available to find the nearest.");
				return null;
			}

			// Display a marker for the current location
			if (this.currentLocationMarker) {
				this.currentLocationMarker.remove();
			}

			const pointsFeatureCollection = turf.featureCollection(
				this.allPoints.map((feature) => ({
					type: "Feature",
					geometry: {
						type: "Point",
						coordinates: feature.coordinates,
					},
					properties: feature.properties,
				}))
			);

			const nearest = turf.nearestPoint(
				turf.point(currentLocation),
				pointsFeatureCollection
			);

			if (nearest) {
				this.showPopupAtPoint(nearest);
			}

			return nearest;
		},

		displayCurrentLocationMarker(currentLocation) {
			// console.log("Current Location:", currentLocation);
			// Remove existing location marker if it exists
			if (this.currentLocationMarker) {
				this.currentLocationMarker.remove();
			}

			// Create a new marker and add it to the map
			this.currentLocationMarker = new mapboxGl.Marker({
				color: "red",
				opacity: 0.75,
			})
				.setLngLat(currentLocation)
				.addTo(this.map);
		},
		// showPopupAtPoint(pointFeature) {
		// 	const { coordinates } = pointFeature.geometry;
		// 	const lngLat = { lng: coordinates[0], lat: coordinates[1] };
		// 	// Assuming `map` is your Mapbox instance and it's already defined in your store
		// 	if (this.map) {
		// 		this.map.flyTo({ center: lngLat, zoom: 15 }); // Optional: fly to the point before showing the popup
		// 		this.addPopup({
		// 			lngLat: lngLat,
		// 			properties: pointFeature.properties, // assuming properties you want to show are here
		// 		});
		// 	}
		// },

		// showPopupAtPoint(pointFeature) {
		// 	const { coordinates } = pointFeature.geometry;
		// 	if (
		// 		!coordinates ||
		// 		coordinates.length < 2 ||
		// 		isNaN(coordinates[0]) ||
		// 		isNaN(coordinates[1])
		// 	) {
		// 		console.error("Invalid coordinates", coordinates);
		// 		return; // 如果坐標無效，則退出函數
		// 	}

		// 	const lngLat = { lng: coordinates[0], lat: coordinates[1] };
		// 	// Assuming `map` is your Mapbox instance and it's already defined in your store
		// 	if (this.map) {
		// 		this.map.flyTo({ center: lngLat, zoom: 15 }); // Optional: fly to the point before showing the popup
		// 		this.addPopup({
		// 			lngLat: lngLat,
		// 			properties: pointFeature.properties, // assuming properties you want to show are here
		// 		});
		// 	}
		// },

		//

		// showPopupAtPoint(pointFeature) {
		// 	if (
		// 		!pointFeature ||
		// 		!pointFeature.geometry ||
		// 		!pointFeature.geometry.coordinates
		// 	) {
		// 		console.error(
		// 			"Invalid point feature or coordinates missing",
		// 			pointFeature
		// 		);
		// 		return;
		// 	}

		// 	const { coordinates } = pointFeature.geometry;
		// 	if (
		// 		coordinates.length < 2 ||
		// 		isNaN(coordinates[0]) ||
		// 		isNaN(coordinates[1])
		// 	) {
		// 		console.error("Invalid coordinates", coordinates);
		// 		return;
		// 	}

		// 	const lngLat = { lng: coordinates[0], lat: coordinates[1] };

		// 	if (this.map) {
		// 		// Remove any existing popup
		// 		if (this.popup) {
		// 			this.popup.remove();
		// 			this.popup = null; // Optional, depends on your popup management logic
		// 		}

		// 		// Optional: Fly to the point
		// 		// this.map.flyTo({ center: lngLat, zoom: 15 });

		// 		// Add a new popup
		// 		this.popup = this.addPopup({
		// 			lngLat: lngLat,
		// 			properties: pointFeature.properties,
		// 		});
		// 	}
		// },

		showPopupAtPoint(pointFeature) {
			if (
				!pointFeature ||
				!pointFeature.geometry ||
				!pointFeature.geometry.coordinates
			) {
				console.error(
					"Invalid point feature or coordinates missing",
					pointFeature
				);
				return;
			}

			const { coordinates } = pointFeature.geometry;
			if (
				coordinates.length < 2 ||
				isNaN(coordinates[0]) ||
				isNaN(coordinates[1])
			) {
				console.error("Invalid coordinates", coordinates);
				return;
			}

			const lngLat = { lng: coordinates[0], lat: coordinates[1] };

			if (this.map) {
				// Remove any existing popup
				if (this.popup) {
					this.popup.remove();
					this.popup = null; // Optional, depends on your popup management logic
				}

				// Optional: Fly to the point
				// this.map.flyTo({ center: lngLat, zoom: 15 });

				// Add a new popup
				this.popup = this.addPopup({
					lngLat: lngLat,
					properties: pointFeature.properties,
				});
			}
		},
		//////
		initializeMapBox() {
			this.map = null;
			const MAPBOXTOKEN = import.meta.env.VITE_MAPBOXTOKEN;
			mapboxGl.accessToken = MAPBOXTOKEN;
			this.map = new mapboxGl.Map({
				...MapObjectConfig,
				style: mapStyle,
			});
			this.map.addControl(new mapboxGl.NavigationControl());
			this.map.doubleClickZoom.disable();
			this.map
				.on("load", () => {
					this.initializeBasicLayers();
				})
				.on("click", (event) => {
					if (this.popup) {
						this.popup = null;
					}
					this.addPopup(event);
				})
				.on("idle", () => {
					this.loadingLayers = this.loadingLayers.filter(
						(el) => el !== "rendering"
					);
				});
		},
		// 2. Adds three basic layers to the map (Taipei District, Taipei Village labels, and Taipei 3D Buildings)
		// Due to performance concerns, Taipei 3D Buildings won't be added in the mobile version
		initializeBasicLayers() {
			const authStore = useAuthStore();
			if (!this.map) return;
			//

			// Taipei District Labels
			fetch(`/mapData/taipei_town.geojson`)
				.then((response) => response.json())
				.then((data) => {
					this.map
						.addSource("taipei_town", {
							type: "geojson",
							data: data,
						})
						.addLayer(TaipeiTown);
				});
			// Taipei Village Labels
			fetch(`/mapData/taipei_village.geojson`)
				.then((response) => response.json())
				.then((data) => {
					this.map
						.addSource("taipei_village", {
							type: "geojson",
							data: data,
						})
						.addLayer(TaipeiVillage);
				});
			// Taipei 3D Buildings
			if (!authStore.isMobileDevice) {
				this.map
					.addSource("taipei_building_3d_source", {
						type: "vector",
						url: import.meta.env.VITE_MAPBOXTILE,
					})
					.addLayer(TaipeiBuilding);
			}
			// Taipei Village Boundaries
			this.map
				.addSource(`tp_village`, {
					type: "vector",
					scheme: "tms",
					tolerance: 0,
					tiles: [
						`${location.origin}/geo_server/gwc/service/tms/1.0.0/taipei_vioc:tp_village@EPSG:900913@pbf/{z}/{x}/{y}.pbf`,
					],
				})
				.addLayer(TpVillage);
			// Taipei District Boundaries
			this.map
				.addSource(`tp_district`, {
					type: "vector",
					scheme: "tms",
					tolerance: 0,
					tiles: [
						`${location.origin}/geo_server/gwc/service/tms/1.0.0/taipei_vioc:tp_district@EPSG:900913@pbf/{z}/{x}/{y}.pbf`,
					],
				})
				.addLayer(TpDistrict);

			this.addSymbolSources();
		},
		// 3. Adds symbols that will be used by some map layers
		addSymbolSources() {
			const images = [
				"metro",
				"triangle_green",
				"triangle_white",
				"bike_green",
				"bike_orange",
				"bike_red",
			];
			images.forEach((element) => {
				this.map.loadImage(
					`/images/map/${element}.png`,
					(error, image) => {
						if (error) throw error;
						this.map.addImage(element, image);
					}
				);
			});
		},
		// 4. Toggle district boundaries
		toggleDistrictBoundaries(status) {
			if (status) {
				this.map.setLayoutProperty(
					"tp_district",
					"visibility",
					"visible"
				);
			} else {
				this.map.setLayoutProperty("tp_district", "visibility", "none");
			}
		},

		// 5. Toggle village boundaries
		toggleVillageBoundaries(status) {
			if (status) {
				this.map.setLayoutProperty(
					"tp_village",
					"visibility",
					"visible"
				);
			} else {
				this.map.setLayoutProperty("tp_village", "visibility", "none");
			}
		},

		/* Adding Map Layers */
		// 1. Passes in the map_config (an Array of Objects) of a component and adds all layers to the map layer list
		addToMapLayerList(map_config) {
			map_config.forEach((element) => {
				let mapLayerId = `${element.index}-${element.type}`;
				// 1-1. If the layer exists, simply turn on the visibility and add it to the visible layers list
				if (
					this.currentLayers.find((element) => element === mapLayerId)
				) {
					this.loadingLayers.push("rendering");
					this.turnOnMapLayerVisibility(mapLayerId);
					if (
						!this.currentVisibleLayers.find(
							(element) => element === mapLayerId
						)
					) {
						this.currentVisibleLayers.push(mapLayerId);
					}
					return;
				}
				let appendLayer = { ...element };
				appendLayer.layerId = mapLayerId;
				// 1-2. If the layer doesn't exist, call an API to get the layer data
				this.loadingLayers.push(appendLayer.layerId);
				if (element.source === "geojson") {
					this.fetchLocalGeoJson(appendLayer);
				} else if (element.source === "raster") {
					this.addRasterSource(appendLayer);
				}
			});
		},
		// 2. Call an API to get the layer data
		// fetchLocalGeoJson(map_config) {
		// 	axios
		// 		.get(`/mapData/${map_config.index}.geojson`)
		// 		.then((rs) => {
		// 			this.addGeojsonSource(map_config, rs.data);
		// 		})
		// 		.catch((e) => console.error(e));
		// },
		fetchLocalGeoJson(map_config) {
			axios
				.get(`/mapData/${map_config.index}.geojson`)
				.then((response) => {
					this.addGeojsonSource(map_config, response.data);
					this.storePoints(response.data, response.data);
				})
				.catch((e) => console.error(e));
		},

		// storePoints(data) {
		// 	const features = data.features.filter(
		// 		(feature) => feature.geometry.type === "circle"
		// 	);
		// 	console.log("Filtered Point Features:", features); // 日誌輸出篩選後的點特徵
		// 	this.allPoints.push(
		// 		...features.map((feature) => ({
		// 			type: feature.type,
		// 			coordinates: feature.geometry.coordinates,
		// 			properties: feature.properties,
		// 			sourceLayerId: sourceLayerId, // 存储来源图层 ID 以便后续操作
		// 		}))
		// 	);
		// },

		// storePoints(data, layerId) {
		// 	const features = data.features.filter(
		// 		(feature) => feature.geometry.type === "Point"
		// 	);
		// 	console.log("Filtered Point Features:", features); // 日誌輸出篩選後的點特徵
		// 	this.allPoints.push(
		// 		...features.map((feature) => {
		// 			const pointData = {
		// 				type: feature.type,
		// 				coordinates: feature.geometry.coordinates,
		// 				properties: feature.properties,
		// 				sourceLayerId: layerId,
		// 			};
		// 			console.log("Storing Point Data:", pointData); // 日誌輸出即將存儲的點數據
		// 			return pointData;
		// 		})
		// 	);
		// },

		storePoints(data, layerId) {
			const features = data.features.filter(
				(feature) => feature.geometry.type === "Point"
			);
			this.allPoints.push(
				...features.map((feature) => ({
					type: feature.type,
					coordinates: feature.geometry.coordinates, // Directly use the coordinates array
					properties: feature.properties,
					sourceLayerId: layerId, // Store source layer ID for reference
				}))
			);
		},

		// 3-1. Add a local geojson as a source in mapbox
		addGeojsonSource(map_config, data) {
			if (!["voronoi", "isoline"].includes(map_config.type)) {
				this.map.addSource(`${map_config.layerId}-source`, {
					type: "geojson",
					data: { ...data },
				});
			}
			if (map_config.type === "arc") {
				this.AddArcMapLayer(map_config, data);
			} else if (map_config.type === "voronoi") {
				this.AddVoronoiMapLayer(map_config, data);
			} else if (map_config.type === "isoline") {
				this.AddIsolineMapLayer(map_config, data);
			} else {
				this.addMapLayer(map_config);
			}
		},
		// 3-2. Add a raster map as a source in mapbox
		addRasterSource(map_config) {
			this.map.addSource(`${map_config.layerId}-source`, {
				type: "vector",
				scheme: "tms",
				tolerance: 0,
				tiles: [
					`${location.origin}/geo_server/gwc/service/tms/1.0.0/taipei_vioc:${map_config.index}@EPSG:900913@pbf/{z}/{x}/{y}.pbf`,
				],
			});
			this.addMapLayer(map_config);
		},
		// 4-1. Using the mapbox source and map config, create a new layer
		// The styles and configs can be edited in /assets/configs/mapbox/mapConfig.js
		addMapLayer(map_config) {
			let extra_paint_configs = {};
			let extra_layout_configs = {};
			if (map_config.icon) {
				extra_paint_configs = {
					...maplayerCommonPaint[
						`${map_config.type}-${map_config.icon}`
					],
				};
				extra_layout_configs = {
					...maplayerCommonLayout[
						`${map_config.type}-${map_config.icon}`
					],
				};
			}
			if (map_config.size) {
				extra_paint_configs = {
					...extra_paint_configs,
					...maplayerCommonPaint[
						`${map_config.type}-${map_config.size}`
					],
				};
				extra_layout_configs = {
					...extra_layout_configs,
					...maplayerCommonLayout[
						`${map_config.type}-${map_config.size}`
					],
				};
			}
			this.loadingLayers.push("rendering");
			this.map.addLayer({
				id: map_config.layerId,
				type: map_config.type,
				"source-layer":
					map_config.source === "raster" ? map_config.index : "",
				paint: {
					...maplayerCommonPaint[`${map_config.type}`],
					...extra_paint_configs,
					...map_config.paint,
				},
				layout: {
					...maplayerCommonLayout[`${map_config.type}`],
					...extra_layout_configs,
				},
				source: `${map_config.layerId}-source`,
			});
			this.currentLayers.push(map_config.layerId);
			this.mapConfigs[map_config.layerId] = map_config;
			this.currentVisibleLayers.push(map_config.layerId);
			this.loadingLayers = this.loadingLayers.filter(
				(el) => el !== map_config.layerId
			);
		},
		// 4-2. Add Map Layer for Arc Maps
		AddArcMapLayer(map_config, data) {
			const authStore = useAuthStore();
			const lines = [...JSON.parse(JSON.stringify(data.features))];
			const arcInterval = 20;

			this.loadingLayers.push("rendering");

			for (let i = 0; i < lines.length; i++) {
				let line = [];
				let lngDif =
					lines[i].geometry.coordinates[1][0] -
					lines[i].geometry.coordinates[0][0];
				let lngInterval = lngDif / arcInterval;
				let latDif =
					lines[i].geometry.coordinates[1][1] -
					lines[i].geometry.coordinates[0][1];
				let latInterval = latDif / arcInterval;

				let maxElevation =
					Math.pow(Math.abs(lngDif * latDif), 0.5) * 80000;

				for (let j = 0; j < arcInterval + 1; j++) {
					let waypointElevation =
						Math.sin((Math.PI * j) / arcInterval) * maxElevation;
					line.push([
						lines[i].geometry.coordinates[0][0] + lngInterval * j,
						lines[i].geometry.coordinates[0][1] + latInterval * j,
						waypointElevation,
					]);
				}

				lines[i].geometry.coordinates = [...line];
			}

			const tb = (window.tb = new Threebox(
				this.map,
				this.map.getCanvas().getContext("webgl"), //get the context from the map canvas
				{ defaultLights: true }
			));

			const delay = authStore.isMobileDevice ? 2000 : 500;

			setTimeout(() => {
				this.map.addLayer({
					id: map_config.layerId,
					type: "custom",
					renderingMode: "3d",
					onAdd: function () {
						const paintSettings = map_config.paint
							? map_config.paint
							: { "arc-color": ["#ffffff"] };
						const gradientSteps = calculateGradientSteps(
							paintSettings["arc-color"][0],
							paintSettings["arc-color"][1]
								? paintSettings["arc-color"][1]
								: paintSettings["arc-color"][0],
							arcInterval + 1
						);
						for (let line of lines) {
							let lineOptions = {
								geometry: line.geometry.coordinates,
								color: 0xffffff,
								width: paintSettings["arc-width"]
									? paintSettings["arc-width"]
									: 2,
								opacity:
									paintSettings["arc-opacity"] ||
									paintSettings["arc-opacity"] === 0
										? paintSettings["arc-opacity"]
										: 0.5,
							};

							let lineMesh = tb.line(lineOptions);
							lineMesh.geometry.setColors(gradientSteps);
							lineMesh.material.vertexColors = true;

							tb.add(lineMesh);
						}
					},
					render: function () {
						tb.update(); //update Threebox scene
					},
				});
				this.currentLayers.push(map_config.layerId);
				this.mapConfigs[map_config.layerId] = map_config;
				this.currentVisibleLayers.push(map_config.layerId);
				this.loadingLayers = this.loadingLayers.filter(
					(el) => el !== map_config.layerId
				);
			}, delay);
		},
		// 4-3. Add Map Layer for Voronoi Maps
		// Developed by 00:21, Taipei Codefest 2023
		AddVoronoiMapLayer(map_config, data) {
			this.loadingLayers.push("rendering");

			let voronoi_source = {
				type: data.type,
				crs: data.crs,
				features: [],
			};

			// Get features alone
			let { features } = data;

			// Get coordnates alone
			let coords = features.map(
				(location) => location.geometry.coordinates
			);

			// Remove duplicate coordinates (so that they wont't cause problems in the Voronoi algorithm...)
			let shouldBeRemoved = coords.map((coord1, ind) => {
				return (
					coords.findIndex((coord2) => {
						return (
							coord2[0] === coord1[0] && coord2[1] === coord1[1]
						);
					}) !== ind
				);
			});

			features = features.filter((_, ind) => !shouldBeRemoved[ind]);
			coords = coords.filter((_, ind) => !shouldBeRemoved[ind]);

			// Calculate cell for each coordinate
			let cells = voronoi(coords);

			// Push cell outlines to source data
			for (let i = 0; i < cells.length; i++) {
				voronoi_source.features.push({
					...features[i],
					geometry: {
						type: "LineString",
						coordinates: cells[i],
					},
				});
			}

			// Add source and layer
			this.map.addSource(`${map_config.layerId}-source`, {
				type: "geojson",
				data: { ...voronoi_source },
			});

			let new_map_config = { ...map_config };
			new_map_config.type = "line";
			this.addMapLayer(new_map_config);
		},
		// 4-4. Add Map Layer for Isoline Maps
		// Developed by 00:21, Taipei Codefest 2023
		AddIsolineMapLayer(map_config, data) {
			this.loadingLayers.push("rendering");
			// Step 1: Generate a 2D scalar field from known data points
			// - Turn the original data into the format that can be accepted by interpolation()
			let dataPoints = data.features.map((item) => {
				return {
					x: item.geometry.coordinates[0],
					y: item.geometry.coordinates[1],
					value: item.properties[
						map_config.paint?.["isoline-key"] || "value"
					],
				};
			});

			let lngStart = 121.42955;
			let lngEnd = 121.68351;
			let latStart = 24.94679;
			let latEnd = 25.21811;

			let targetPoints = [];
			let gridSize = 0.001;
			let rowN = 0;
			let colN = 0;

			// - Generate target point coordinates
			for (let i = latStart; i <= latEnd; i += gridSize, rowN += 1) {
				colN = 0;
				for (let j = lngStart; j <= lngEnd; j += gridSize, colN += 1) {
					targetPoints.push({ x: j, y: i });
				}
			}

			// - Get target points interpolation result
			let interpolationResult = interpolation(dataPoints, targetPoints);

			// Step 2: Calculate isolines from the 2D scalar field
			// - Turn the interpolation result into the format that can be accepted by marchingSquare()
			let discreteData = [];
			for (let y = 0; y < rowN; y++) {
				discreteData.push([]);
				for (let x = 0; x < colN; x++) {
					discreteData[y].push(interpolationResult[y * colN + x]);
				}
			}

			// - Initialize geojson data
			let isoline_data = {
				type: "FeatureCollection",
				crs: {
					type: "name",
					properties: { name: "urn:ogc:def:crs:OGC:1.3:CRS84" },
				},
				features: [],
			};

			// - Repeat the marching square algorithm for differnt iso-values (40, 42, 44 ... 74 in this case)
			for (let isoValue = 40; isoValue <= 75; isoValue += 2) {
				let result = marchingSquare(discreteData, isoValue);

				let transformedResult = result.map((line) => {
					return line.map((point) => {
						return [
							point[0] * gridSize + lngStart,
							point[1] * gridSize + latStart,
						];
					});
				});

				isoline_data.features = isoline_data.features.concat(
					// Turn result into geojson format
					transformedResult.map((line) => {
						return {
							type: "Feature",
							properties: { value: isoValue },
							geometry: { type: "LineString", coordinates: line },
						};
					})
				);
			}

			// Step 3: Add source and layer
			this.map.addSource(`${map_config.layerId}-source`, {
				type: "geojson",

				data: { ...isoline_data },
			});

			delete map_config.paint?.["isoline-key"];

			let new_map_config = { ...map_config, type: "line" };
			this.addMapLayer(new_map_config);
		},
		//  5. Turn on the visibility for a exisiting map layer
		turnOnMapLayerVisibility(mapLayerId) {
			this.map.setLayoutProperty(mapLayerId, "visibility", "visible");
			const layerData = this.map.getSource(`${mapLayerId}-source`)._data;
			if (layerData && layerData.features) {
				this.storePoints(layerData);
			}
		},
		// 6. Turn off the visibility of an exisiting map layer but don't remove it completely
		turnOffMapLayerVisibility(map_config) {
			map_config.forEach((element) => {
				let mapLayerId = `${element.index}-${element.type}`;
				this.loadingLayers = this.loadingLayers.filter(
					(el) => el !== mapLayerId
				);

				if (this.map.getLayer(mapLayerId)) {
					this.map.setFilter(mapLayerId, null);
					this.map.setLayoutProperty(
						mapLayerId,
						"visibility",
						"none"
					);
				}
				this.currentVisibleLayers = this.currentVisibleLayers.filter(
					(element) => element !== mapLayerId
				);

				this.allPoints = [];

				this.removePopup();
			});
		},

		/* Popup Related Functions */
		// 1. Adds a popup when the user clicks on a item. The event will be passed in.
		addPopup(event) {
			// Gets the info that is contained in the coordinates that the user clicked on (only visible layers)

			const clickFeatureDatas = this.map.queryRenderedFeatures(
				event.point,
				{
					layers: this.currentVisibleLayers,
				}
			);
			// Return if there is no info in the click
			if (!clickFeatureDatas || clickFeatureDatas.length === 0) {
				return;
			}

			if (this.popup) {
				this.popup.remove();
				this.popup = null;
			}

			// Parse clickFeatureDatas to get the first 3 unique layer datas, skip over already included layers
			const mapConfigs = [];
			const parsedPopupContent = [];
			let previousParsedLayer = "";

			for (let i = 0; i < clickFeatureDatas.length; i++) {
				if (mapConfigs.length === 3) break;
				if (previousParsedLayer === clickFeatureDatas[i].layer.id)
					continue;
				previousParsedLayer = clickFeatureDatas[i].layer.id;
				mapConfigs.push(this.mapConfigs[clickFeatureDatas[i].layer.id]);
				parsedPopupContent.push(clickFeatureDatas[i]);
			}
			// Create a new mapbox popup
			this.popup = new mapboxGl.Popup()
				.setLngLat(event.lngLat)
				.setHTML('<div id="vue-popup-content"></div>')
				.addTo(this.map);
			// Mount a vue component (MapPopup) to the id "vue-popup-content" and pass in data
			const PopupComponent = defineComponent({
				extends: MapPopup,
				setup() {
					// Only show the data of the topmost layer
					return {
						popupContent: parsedPopupContent,
						mapConfigs: mapConfigs,
						activeTab: ref(0),
					};
				},
			});
			// This helps vue determine the most optimal time to mount the component
			nextTick(() => {
				const app = createApp(PopupComponent);
				app.mount("#vue-popup-content");
			});
		},
		// 2. Remove the current popup
		removePopup() {
			if (this.popup) {
				this.popup.remove();
			}
			this.popup = null;
		},

		/* Functions that change the viewing experience of the map */
		// 1. Add new saved location that users can quickly zoom to
		addNewSavedLocation(name) {
			const coordinates = this.map.getCenter();
			const zoom = this.map.getZoom();
			const pitch = this.map.getPitch();
			const bearing = this.map.getBearing();
			this.savedLocations.push([coordinates, zoom, pitch, bearing, name]);
		},
		// 2. Zoom to a location
		// [[lng, lat], zoom, pitch, bearing, savedLocationName]
		easeToLocation(location_array) {
			this.map.easeTo({
				center: location_array[0],
				zoom: location_array[1],
				duration: 4000,
				pitch: location_array[2],
				bearing: location_array[3],
			});
		},
		// 3. Fly to a location
		flyToLocation(location_array) {
			this.map.flyTo({
				center: location_array,
				duration: 1000,
			});
		},
		// 4. Remove a saved location
		removeSavedLocation(index) {
			this.savedLocations.splice(index, 1);
		},
		// 5. Force map to resize after sidebar collapses
		resizeMap() {
			if (this.map) {
				setTimeout(() => {
					this.map.resize();
				}, 200);
			}
		},

		/* Map Filtering */
		// 1. Add a filter based on a each map layer's properties (byParam)
		filterByParam(map_filter, map_configs, xParam, yParam) {
			// If there are layers loading, don't filter
			if (this.loadingLayers.length > 0) return;
			const dialogStore = useDialogStore();
			if (!this.map || dialogStore.dialogs.moreInfo) {
				return;
			}
			map_configs.map((map_config) => {
				let mapLayerId = `${map_config.index}-${map_config.type}`;
				if (map_config && map_config.type === "arc") {
					// Only turn off original layer visibility
					this.map.setLayoutProperty(
						mapLayerId,
						"visibility",
						"none"
					);
					// Remove any existing filtered layer
					if (this.map.getLayer(`${mapLayerId}-filtered`)) {
						this.map.removeLayer(`${mapLayerId}-filtered`);
					}
					// Filter data to render new filtered layer
					let toBeFiltered = {
						...this.map.getSource(`${mapLayerId}-source`)._data,
					};
					if (xParam) {
						toBeFiltered.features = toBeFiltered.features.filter(
							(el) =>
								el.properties[map_filter.byParam.xParam] ===
								xParam
						);
					}
					if (yParam) {
						toBeFiltered.features = toBeFiltered.features.filter(
							(el) =>
								el.properties[map_filter.byParam.yParam] ===
								yParam
						);
					}
					map_config.layerId = `${mapLayerId}-filtered`;
					// Add new filtered layer
					this.AddArcMapLayer(map_config, toBeFiltered);
					return;
				}
				// If x and y both exist, filter by both
				if (
					map_filter.byParam.xParam &&
					map_filter.byParam.yParam &&
					xParam &&
					yParam
				) {
					this.map.setFilter(mapLayerId, [
						"all",
						["==", ["get", map_filter.byParam.xParam], xParam],
						["==", ["get", map_filter.byParam.yParam], yParam],
					]);
				}
				// If only y exists, filter by y
				else if (map_filter.byParam.yParam && yParam) {
					this.map.setFilter(mapLayerId, [
						"==",
						["get", map_filter.byParam.yParam],
						yParam,
					]);
				}
				// default to filter by x
				else if (map_filter.byParam.xParam && xParam) {
					this.map.setFilter(mapLayerId, [
						"==",
						["get", map_filter.byParam.xParam],
						xParam,
					]);
				}
			});
		},
		// 2. filter by layer name (byLayer)
		filterByLayer(map_configs, xParam) {
			const dialogStore = useDialogStore();
			// If there are layers loading, don't filter
			if (this.loadingLayers.length > 0) return;
			if (!this.map || dialogStore.dialogs.moreInfo) {
				return;
			}
			map_configs.map((map_config) => {
				let mapLayerId = `${map_config.index}-${map_config.type}`;
				if (map_config.title !== xParam) {
					this.map.setLayoutProperty(
						mapLayerId,
						"visibility",
						"none"
					);
				} else {
					this.map.setLayoutProperty(
						mapLayerId,
						"visibility",
						"visible"
					);
				}
			});
		},
		// 3. Remove any property filters on a map layer
		clearByParamFilter(map_configs) {
			const dialogStore = useDialogStore();
			if (!this.map || dialogStore.dialogs.moreInfo) {
				return;
			}
			map_configs.map((map_config) => {
				let mapLayerId = `${map_config.index}-${map_config.type}`;
				if (map_config && map_config.type === "arc") {
					if (this.map.getLayer(`${mapLayerId}-filtered`)) {
						this.map.removeLayer(`${mapLayerId}-filtered`);
					}
					this.currentLayers = this.currentLayers.filter(
						(item) => item !== `${mapLayerId}-filtered`
					);
					this.currentVisibleLayers =
						this.currentVisibleLayers.filter(
							(item) => item !== `${mapLayerId}-filtered`
						);
					this.map.setLayoutProperty(
						mapLayerId,
						"visibility",
						"visible"
					);
					return;
				}
				this.map.setFilter(mapLayerId, null);
			});
		},
		// 4. Remove any layer filters on a map layer.
		clearByLayerFilter(map_configs) {
			const dialogStore = useDialogStore();
			if (!this.map || dialogStore.dialogs.moreInfo) {
				return;
			}
			map_configs.map((map_config) => {
				let mapLayerId = `${map_config.index}-${map_config.type}`;
				this.map.setLayoutProperty(mapLayerId, "visibility", "visible");
			});
		},

		/* Clearing the map */
		// 1. Called when the user is switching between maps
		clearOnlyLayers() {
			this.currentLayers.forEach((element) => {
				this.map.removeLayer(element);
				if (this.map.getSource(`${element}-source`)) {
					this.map.removeSource(`${element}-source`);
				}
			});
			this.currentLayers = [];
			this.mapConfigs = {};
			this.currentVisibleLayers = [];
			this.removePopup();
		},
		// 2. Called when user navigates away from the map
		clearEntireMap() {
			this.currentLayers = [];
			this.mapConfigs = {};
			this.map = null;
			this.currentVisibleLayers = [];
			this.removePopup();
		},
	},
});
