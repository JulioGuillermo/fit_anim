// @ts-check
import { initWebAssembly } from "./app_wasm.js";
import { initScene } from "./three.js";

// function updateAnim() {
//   if ("updateHTML" in globalThis) {
//     globalThis.updateHTML();
//   }
//   requestAnimationFrame(updateAnim);
// }
// requestAnimationFrame(updateAnim);

async function initApp() {
  await initWebAssembly();
  await initScene();
}
initApp();
