// @ts-check
import * as THREE from "three";
import { EffectComposer } from "three/addons/postprocessing/EffectComposer.js";
import { RenderPass } from "three/addons/postprocessing/RenderPass.js";
import { OutputPass } from "three/addons/postprocessing/OutputPass.js";
import { BloomPass } from "three/addons/postprocessing/BloomPass.js";

globalThis.THREE = THREE;

export async function initScene() {
  const scene = new THREE.Scene();
  const camera = new THREE.PerspectiveCamera(
    75,
    window.innerWidth / window.innerHeight,
    0.1,
    1000,
  );
  // new THREE.MeshPhysicalMaterial().emissiveIntensity

  const renderer = new THREE.WebGLRenderer({
    antialias: true,
  });
  renderer.setSize(window.innerWidth, window.innerHeight);
  document.body.appendChild(renderer.domElement);
  const composer = new EffectComposer(renderer);

  const renderPass = new RenderPass(scene, camera);
  composer.addPass(renderPass);

  const outputPass = new OutputPass();
  composer.addPass(outputPass);

  if ("goInitScene" in globalThis) {
    globalThis.goInitScene(scene, camera, renderer);
  }

  function animationFrame() {
    if ("goRender" in globalThis) {
      globalThis.goRender(scene, camera, renderer);
    }
    composer.render();
    // renderer.render(scene, camera);

    requestAnimationFrame(animationFrame);
  }
  // renderer.setAnimationLoop(() => {
  // });
  animationFrame();
}
