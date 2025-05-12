import { useQuery } from '@tanstack/react-query'
import api from '../shared/api'

export function useLists() {
  return useQuery({
    queryKey: ['lists'],
    queryFn: () =>
      api.get('/api/lists')
         .then(res => res.data.data ?? []),
  })
}