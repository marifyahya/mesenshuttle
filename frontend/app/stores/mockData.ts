import { defineStore } from 'pinia'

export const useMockDataStore = defineStore('mockData', {
  state: () => ({
    routes: [
      { id: 1, originCity: 'Jakarta', originPool: 'Lebak Bulus', destinationCity: 'Bandung', destinationPool: 'Cihampelas' },
      { id: 2, originCity: 'Jakarta', originPool: 'Kelapa Gading', destinationCity: 'Bandung', destinationPool: 'Miko Mall' },
      { id: 3, originCity: 'Bandung', originPool: 'Buah Batu', destinationCity: 'Yogyakarta', destinationPool: 'Giwangan' },
      { id: 4, originCity: 'Bandung', originPool: 'Cihampelas', destinationCity: 'Jakarta', destinationPool: 'Lebak Bulus' },
      { id: 5, originCity: 'Cirebon', originPool: 'Harjamukti', destinationCity: 'Bandung', destinationPool: 'Cihampelas' },
      { id: 6, originCity: 'Bandung', originPool: 'Miko Mall', destinationCity: 'Cirebon', destinationPool: 'Harjamukti' },
      { id: 7, originCity: 'Cirebon', originPool: 'Harjamukti', destinationCity: 'Jakarta', destinationPool: 'Kelapa Gading' },
      { id: 8, originCity: 'Jakarta', originPool: 'Lebak Bulus', destinationCity: 'Cirebon', destinationPool: 'Harjamukti' }
    ],
    fleets: [
      { id: 1, name: 'MesenShuttle 05', vehicle: 'Shuttle', type: 'Reguler', capacity: 12 },
      { id: 2, name: 'MesenShuttle 01', vehicle: 'Shuttle', type: 'Premium', capacity: 10 },
      { id: 3, name: 'MesenShuttle 02', vehicle: 'Shuttle', type: 'Reguler', capacity: 12 },
      { id: 4, name: 'MesenShuttle 06', vehicle: 'Shuttle', type: 'Premium', capacity: 10 },
      { id: 5, name: 'MesenShuttle 07', vehicle: 'Shuttle', type: 'Reguler', capacity: 12 },
      { id: 6, name: 'MesenShuttle 03', vehicle: 'Shuttle', type: 'Premium', capacity: 10 },
      { id: 7, name: 'MesenShuttle 04', vehicle: 'Shuttle', type: 'Reguler', capacity: 12 }
    ],
    schedules: [
      // Jakarta -> Bandung
      { id: 1, routeId: 1, fleetId: 1, departureTime: '2026-07-10T08:00:00Z', price: 150000 },
      { id: 2, routeId: 1, fleetId: 2, departureTime: '2026-07-10T09:00:00Z', price: 180000 },
      { id: 3, routeId: 1, fleetId: 3, departureTime: '2026-07-10T10:00:00Z', price: 120000 },
      { id: 4, routeId: 1, fleetId: 4, departureTime: '2026-07-10T13:00:00Z', price: 200000 },
      { id: 5, routeId: 1, fleetId: 6, departureTime: '2026-07-10T15:30:00Z', price: 195000 },

      // Bandung -> Jakarta
      { id: 6, routeId: 4, fleetId: 1, departureTime: '2026-07-10T07:30:00Z', price: 150000 },
      { id: 7, routeId: 4, fleetId: 2, departureTime: '2026-07-10T10:30:00Z', price: 180000 },
      { id: 8, routeId: 4, fleetId: 7, departureTime: '2026-07-10T14:00:00Z', price: 120000 },
      { id: 9, routeId: 4, fleetId: 5, departureTime: '2026-07-10T18:00:00Z', price: 145000 },

      // Cirebon -> Bandung
      { id: 10, routeId: 5, fleetId: 3, departureTime: '2026-07-10T06:00:00Z', price: 95000 },
      { id: 11, routeId: 5, fleetId: 6, departureTime: '2026-07-10T11:00:00Z', price: 140000 },
      { id: 12, routeId: 5, fleetId: 4, departureTime: '2026-07-10T16:00:00Z', price: 160000 },
      { id: 13, routeId: 5, fleetId: 7, departureTime: '2026-07-10T19:30:00Z', price: 95000 },

      // Bandung -> Cirebon
      { id: 14, routeId: 6, fleetId: 3, departureTime: '2026-07-10T08:00:00Z', price: 95000 },
      { id: 15, routeId: 6, fleetId: 6, departureTime: '2026-07-10T13:30:00Z', price: 140000 },
      { id: 16, routeId: 6, fleetId: 1, departureTime: '2026-07-10T17:00:00Z', price: 120000 },

      // Cirebon -> Jakarta
      { id: 17, routeId: 7, fleetId: 5, departureTime: '2026-07-10T05:30:00Z', price: 180000 },
      { id: 18, routeId: 7, fleetId: 2, departureTime: '2026-07-10T09:00:00Z', price: 230000 },
      { id: 19, routeId: 7, fleetId: 7, departureTime: '2026-07-10T14:30:00Z', price: 190000 },
      { id: 20, routeId: 7, fleetId: 4, departureTime: '2026-07-10T22:00:00Z', price: 250000 },

      // Jakarta -> Cirebon
      { id: 21, routeId: 8, fleetId: 5, departureTime: '2026-07-10T06:30:00Z', price: 180000 },
      { id: 22, routeId: 8, fleetId: 2, departureTime: '2026-07-10T11:00:00Z', price: 230000 },
      { id: 23, routeId: 8, fleetId: 7, departureTime: '2026-07-10T16:45:00Z', price: 190000 },
      { id: 24, routeId: 8, fleetId: 4, departureTime: '2026-07-10T20:00:00Z', price: 250000 },
    ]
  }),
  getters: {
    getScheduleDetails: (state) => {
      return (scheduleId: number) => {
        const schedule = state.schedules.find(s => s.id === scheduleId)
        if (!schedule) return null;

        const route = state.routes.find(r => r.id === schedule.routeId)
        const fleet = state.fleets.find(f => f.id === schedule.fleetId)

        // Mock remaining seats deterministically to avoid hydration mismatch
        const bookedSeatsCount = (scheduleId * 3) % (fleet?.capacity || 20)
        const availableSeats = (fleet?.capacity || 20) - bookedSeatsCount

        return { ...schedule, route, fleet, availableSeats }
      }
    }
  },
  actions: {
    // Routes
    addRoute(route: any) {
      const id = Math.max(0, ...this.routes.map(r => r.id)) + 1
      this.routes.push({ ...route, id })
    },
    updateRoute(id: number, route: any) {
      const index = this.routes.findIndex(r => r.id === id)
      if (index !== -1) this.routes[index] = { ...this.routes[index], ...route }
    },
    deleteRoute(id: number) {
      this.routes = this.routes.filter(r => r.id !== id)
      // Cascade delete schedules for this route
      this.schedules = this.schedules.filter(s => s.routeId !== id)
    },
    // Fleets
    addFleet(fleet: any) {
      const id = Math.max(0, ...this.fleets.map(f => f.id)) + 1
      this.fleets.push({ ...fleet, id })
    },
    updateFleet(id: number, fleet: any) {
      const index = this.fleets.findIndex(f => f.id === id)
      if (index !== -1) this.fleets[index] = { ...this.fleets[index], ...fleet }
    },
    deleteFleet(id: number) {
      this.fleets = this.fleets.filter(f => f.id !== id)
      // Cascade delete schedules for this fleet
      this.schedules = this.schedules.filter(s => s.fleetId !== id)
    },
    // Schedules
    addSchedule(schedule: any) {
      const id = Math.max(0, ...this.schedules.map(s => s.id)) + 1
      this.schedules.push({ ...schedule, id })
    },
    updateSchedule(id: number, schedule: any) {
      const index = this.schedules.findIndex(s => s.id === id)
      if (index !== -1) this.schedules[index] = { ...this.schedules[index], ...schedule }
    },
    deleteSchedule(id: number) {
      this.schedules = this.schedules.filter(s => s.id !== id)
    }
  }
})
