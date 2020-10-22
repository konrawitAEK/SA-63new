import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import adddata from './components/Adddata';
import main from './components/Main';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/user', adddata);
    router.registerRoute('/main', main);
  },
});
