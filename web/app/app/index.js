import { initDevTools } from 'inferno-devtools';
import { render } from 'inferno';
import App from 'App';

initDevTools();

render(<App />, document.getElementById('app'));
