import React, { useContext, useMemo } from "react";
import styles from "./meetCard.module.css";
import { LineSeparatorMeetCard } from "../icons/icons";
import { WebsocketContext } from "../../contexts/websocket_provider";

type MeetCardProps = {
	chatId: number;
	title: string;
	setActive: (active: boolean) => void;
	userId: string;
	username: string;
	joinChat: (chatId: string) => void;
};

export const MeetCard = (props: MeetCardProps) => {
	const { chatId, title, userId, username, setActive, joinChat } = props;
	const { setConn } = useContext(WebsocketContext);

	return (
		<div className={styles.wrapper__meetCard}>
			<div className={styles.wrapper__users}></div>
			<div className={styles.wrapper__separator}>
				<LineSeparatorMeetCard />
			</div>
			<div className={styles.wrapper__text}>
				<h1 className={styles.text}>{title}</h1>
			</div>
			<div className={styles.wrapper__separator}>
				<LineSeparatorMeetCard />
			</div>
			<div className={styles.wrapper__activities}>
				<h2>{chatId}</h2>
			</div>
			<div>
				<button
					onClick={(event) => {
						event.preventDefault();
						joinChat(String(chatId));
					}}
				>
					ЗАХОДИ!!!
				</button>
			</div>
		</div>
	);
};
