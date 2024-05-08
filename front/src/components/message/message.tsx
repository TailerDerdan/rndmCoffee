import React from "react";
import { Message } from "../chat/chat";
import styles from "./message.module.css";
import { useLocalStorage } from "../../hooks/useLocalStorage";

type MessageBodyProps = {
	messages: Array<Message>;
};

export const MessageBody = (props: MessageBodyProps) => {
	const { messages } = props;

	const [id_user] = useLocalStorage({
		initialValue: "",
		key: "id_user",
	});

	return (
		<>
			{messages.map((message: Message, index: number) => {
				if (message.user_id === String(id_user)) {
					return (
						<div
							key={index}
							className={styles.wrapper__messageSelf}
						>
							<div className={styles.wrapper__username}>
								{message.username}
							</div>
							<div className={styles.wrapper__description}>
								{message.description}
							</div>
						</div>
					);
				} else {
					return (
						<div
							key={index}
							className={styles.wrapper__messageRecv}
						>
							<div className={styles.wrapper__username}>
								{message.username}
							</div>
							<div className={styles.wrapper__description}>
								{message.description}
							</div>
						</div>
					);
				}
			})}
		</>
	);
};
