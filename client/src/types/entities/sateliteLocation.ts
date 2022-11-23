export interface SateliteLocation {
  id_key: number;
  name: string;
  latitude: number;
  longitude: number;
  altitude: number;
  velocity: number;
  visibility: number;
  footprint: number;
  timestamp: number;
  daynum: number;
  solar_lat: number;
  solar_lon: number;
  units: string;
  timezone: string;
  date: string;
  hour: string;
  date_iso: Date;
}
