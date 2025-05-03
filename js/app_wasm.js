// @ts-check

export async function initWebAssembly() {
  if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
      const source = await (await resp).arrayBuffer();
      return await WebAssembly.instantiate(source, importObject);
    };
  }
  const go = new globalThis.Go();
  const result = await WebAssembly.instantiateStreaming(
    fetch("/fit_anim/app.wasm"),
    go.importObject,
  );
  const mod = result.module;
  const inst = result.instance;
  go.run(inst);
  // inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
}
