// @ts-check
import * as THREE from "three";
import { GLTFLoader } from "three/addons/loaders/GLTFLoader.js";

/** @param {THREE.Scene} scene */
export async function loadScene(scene) {
  const loader = new GLTFLoader();
  loader.load(
    "/fit_anim/models/scene4.glb",
    function (gltf) {
      scene.add(gltf.scene);
    },
    undefined,
    function (error) {
      console.error(error);
    },
  );
}
