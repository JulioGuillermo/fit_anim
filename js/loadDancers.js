// @ts-check
import * as THREE from "three";
import { FBXLoader } from "three/addons/loaders/FBXLoader.js";
import { ColladaLoader } from "three/addons/loaders/ColladaLoader.js";

/** @type {THREE.AnimationMixer} */
var mixer;
var mixerL;
var mixerR;

const clock = new THREE.Clock();

/** @param {THREE.Scene} scene */
export async function loadDancers(scene) {
  let loader = new FBXLoader();
  loader.load("/fit_anim/models/SambaDancing.fbx", function (obj) {
    if (obj.animations && obj.animations.length) {
      mixer = new THREE.AnimationMixer(obj);
      const action = mixer.clipAction(obj.animations[0]);
      action.play();
    }
    obj.scale.x = 0.007;
    obj.scale.y = 0.007;
    obj.scale.z = 0.007;
    obj.position.y = -0.39;
    obj.position.z = -0;
    scene.add(obj);
  });

  loader = new ColladaLoader();
  loader.load("/fit_anim/models/stormtrooper.dae", function (collada) {
    const avatar = collada.scene;
    avatar.name = "DancerL";

    avatar.scale.x = 0.07;
    avatar.scale.y = 0.07;
    avatar.scale.z = 0.07;

    avatar.position.z = 0.6;
    avatar.position.y = 0.07;
    avatar.position.x = -1.3;

    avatar.rotation.z = -90;

    const animations = avatar.animations;

    mixerL = new THREE.AnimationMixer(avatar);
    mixerL.clipAction(animations[0]).play();

    scene.add(avatar);
  });
  loader.load("/fit_anim/models/stormtrooper.dae", function (collada) {
    const avatar = collada.scene;
    avatar.name = "DancerR";

    avatar.scale.x = 0.07;
    avatar.scale.y = 0.07;
    avatar.scale.z = 0.07;

    avatar.position.z = 0.6;
    avatar.position.y = 0.07;
    avatar.position.x = 1.3;

    avatar.rotation.z = 90;

    const animations = avatar.animations;

    mixerR = new THREE.AnimationMixer(avatar);
    mixerR.clipAction(animations[0]).play();

    scene.add(avatar);
  });
}

export function updateAnimations() {
  const delta = clock.getDelta();
  if (mixer) mixer.update(delta);
  if (mixerL) mixerL.update(delta);
  if (mixerR) mixerR.update(delta);
}
