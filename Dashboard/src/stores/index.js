import { createPinia } from 'pinia';
import { createPersistedState } from 'pinia-plugin-persistedstate';

export const stores = createPinia();
stores.use(createPersistedState());

export * from './user';

export default stores;