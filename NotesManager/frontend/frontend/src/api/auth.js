import api from '@/shared/api'

export const signIn = async credentials => {
  console.log('→ signIn payload:', credentials)
  const res = await api.post('/auth/sign-in', credentials)
  return res.data.token ?? res.data.accessToken
}
export const signUp = body =>
    api.post('/auth/sign-up', body).then(r => r.data.id)
