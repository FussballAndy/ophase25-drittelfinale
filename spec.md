# API

Base: `/api`

Routes:
- `POST /token`: Verifies Tutor Token
- `POST /winner AUTH tutor`: Send the winner of the current round
- `GET /stations`: Get all stations
- `GET /groups`: Get a list of all groups. Optionally `?short` for points only
- `WS /drittel`: WebSocket Stream for Drittelfinale
- `WS /admin`: WebSocket Stream for Admin and Presentation

# Pages

Routes:
- `GET /`: Base page
- `GET /results`: Table Listing current results