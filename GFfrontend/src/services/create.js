import { request } from 'umi';

export async function createGroup(params) {
  return request('/api/createGroup', {
    data: params,
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  });
}

export async function createPost(params) {
  return request('/api/createPost', {
    data: params,
    method: 'POST',
  });
}