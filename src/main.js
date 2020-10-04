import App from './App.svelte';
import {deserializeConfig, serializeConfig} from "./utils";
import {configStore} from "./store";

const app = new App({
  target: document.body,
  props: {
    name: 'world',
  },
});

let configLoaded = false;

// when config from store changes
configStore.subscribe(updatedConfig => {
  console.log({updatedConfig})
  if (!configLoaded) {
    return;
  }
  // update URL hash after store value has been changed
  const serializedConfig = serializeConfig(updatedConfig);
  window.location.hash = "#" + serializedConfig;
})

// on main app start, parse state from URL hash
const hash = window.location.hash;
if (hash && hash.length > 1) {
  const serializedConfig = hash.substr(1);
  const config = deserializeConfig(serializedConfig);
  configStore.set(config);
  configLoaded = true;
} else {
  configLoaded = true;
}

export default app;
