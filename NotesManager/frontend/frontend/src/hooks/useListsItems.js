import { useQuery } from '@tanstack/react-query'
import api from '../shared/api'
export function useListItems(listId) {
  return useQuery({
    queryKey: ['items', listId],
    queryFn: async () => {
      const res = await api.get(`/api/lists/${listId}/items`)
      return res.data.data
    },
    enabled: !!listId,
  })
}
