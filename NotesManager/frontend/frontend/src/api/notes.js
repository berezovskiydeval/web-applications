import axios from 'axios';

const api = axios.create({ baseURL: '/api', withCredentials: true });

export const getLists   = () => api.get('/lists').then(r => r.data.data);
export const createList = (title) => api.post('/lists', { title });
export const getItems   = (listId) =>
  api.get(`/lists/${listId}/items`).then(r => r.data.data);
export const createItem = (listId, body) =>
  api.post(`/lists/${listId}/items`, body);
