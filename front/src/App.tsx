import React from "react";
import { Main } from "./views/app/main";
import WebSocketProvider from "./contexts/websocket_provider";

function App() {
	return (
		<WebSocketProvider>
			<Main />
		</WebSocketProvider>
	);
}

export default App;
