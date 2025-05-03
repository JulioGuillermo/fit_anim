// @ts-check
import * as THREE from "three";
import { loadScene } from "./load.js";
import {
  BloomEffect,
  EffectComposer,
  EffectPass,
  RenderPass,
} from "postprocessing";
import { loadDancers, updateAnimations } from "./loadDancers.js";

import { WebGLPathTracer } from "three-gpu-pathtracer";

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
  await loadDancers(scene);

  const renderer = new THREE.WebGLRenderer({
    antialias: true,
  });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.toneMapping = THREE.ACESFilmicToneMapping;
  document.body.appendChild(renderer.domElement);

  // const pathTracer = new WebGLPathTracer(renderer);
  // pathTracer.setScene(scene, camera);

  const composer = new EffectComposer(renderer);
  composer.addPass(new RenderPass(scene, camera));
  const bloomEffect = new BloomEffect({
    mipmapBlur: true,
    luminanceThreshold: 0.1,
    radius: 0.5,
    intensity: 5,
  });
  composer.addPass(new EffectPass(camera, bloomEffect));

  if ("goInitScene" in globalThis) {
    globalThis.goInitScene(scene, camera, renderer);
  }

  function animationFrame() {
    if ("goRender" in globalThis) {
      globalThis.goRender(scene, camera, renderer);
    }
    updateAnimations();
    camera.aspect = window.innerWidth / window.innerHeight;
    renderer.setSize(window.innerWidth, window.innerHeight);
    composer.render();
    // pathTracer.renderSample();
  }
  renderer.setAnimationLoop(animationFrame);
}
