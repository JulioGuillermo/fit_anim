// @ts-check
import * as THREE from "three";
import { loadScene } from "./load.js";

globalThis.THREE = THREE;

export async function initScene() {
  const scene = new THREE.Scene();
  const camera = new THREE.PerspectiveCamera(
    75,
    window.innerWidth / window.innerHeight,
    0.1,
    1000,
  );

  await loadScene(scene);

  const renderer = new THREE.WebGLRenderer({
    antialias: true,
  });
  renderer.setSize(window.innerWidth, window.innerHeight);
  document.body.appendChild(renderer.domElement);

  if ("goInitScene" in globalThis) {
    globalThis.goInitScene(scene, camera, renderer);
  }

  function animationFrame() {
    if ("goRender" in globalThis) {
      globalThis.goRender(scene, camera, renderer);
    }
    renderer.render(scene, camera);

    requestAnimationFrame(animationFrame);
  }
  // renderer.setAnimationLoop(() => {
  // });
  animationFrame();
}
