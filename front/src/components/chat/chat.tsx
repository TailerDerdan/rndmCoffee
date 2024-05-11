import React, { useContext, useEffect, useRef, useState } from "react";
import styles from "./chat.module.css";
import { useLocalStorage } from "../../hooks/useLocalStorage";
import { Button, ButtonType } from "../button/button";
import { SendMessageIcon } from "../icons/icons";
import { MessageBody } from "../message/message";

type ChatProps = {
	setActive: (active: boolean) => void;
	active: boolean;
	conn: WebSocket | null;
};

export type Message = {
	id: number;
	description: string;
	user_id: string;
	username: string;
	chatlist_id: string;
};

export const Chat = (props: ChatProps) => {
	const { setActive, active, conn } = props;
	const [messages, setMessages] = useState<Array<Message>>([]);
	const [users, setUsers] = useState<Array<{ username: string }>>([]);
	const textarea = useRef<HTMLTextAreaElement>(null);
	const refMes = useRef<HTMLDivElement>(null);
	const [message, setMessage] = useState("");

	const [token] = useLocalStorage({
		initialValue: "",
		key: "token",
	});

	const [id_user] = useLocalStorage({
		initialValue: "",
		key: "id_user",
	});

	useEffect(() => {
		if (conn === undefined || conn === null) {
			console.log("соединение утратилось( при получении пользователей");
			setActive(false);
			return;
		}

		const roomId = conn.url.split("/")[5];
		async function getUsers() {
			try {
				const res = await fetch(
					`http://localhost:8000/ws/getClients/${roomId}`,
					{
						method: "GET",
						headers: {
							"Content-Type": "application/json",
							Authorization: `Bearer ${token}`,
						},
						credentials: "include",
					},
				);

				const data = await res.json();

				setUsers(data);
			} catch (err) {
				console.error(err);
			}
		}
		getUsers();
	}, []);

	useEffect(() => {
		if (textarea.current) {
			textarea.current!.style.height = "auto";
			textarea.current!.style.height = `${
				textarea.current!.scrollHeight
			}px`;
		}
	}, [message]);

	useEffect(() => {
		if (conn === undefined || conn === null) {
			console.log("соединение утратилось(");
			setActive(false);
			return;
		}

		const roomId = conn.url.split("/")[5];

		conn.onmessage = (message) => {
			const m: Message = JSON.parse(message.data);
			console.log(m);
			if (m.description === "Пользователь зашел в комнату") {
				setUsers([...users, { username: m.username }]);
				return;
			}

			if (m.description === "Пользователь покинул чат") {
				const deleteUser = users.filter(
					(user) => user.username != m.username,
				);
				setUsers([...deleteUser]);
				setMessages([...messages, m]);
				return;
			}

			const saveMessage = async () => {
				const dataMes = {
					user_id: String(m.user_id),
					username: m.username,
					description: m.description,
					chatlist_id: m.chatlist_id,
				};
				console.log(dataMes);
				const res = await fetch(
					`http://localhost:8000/api/chats/${roomId[0]}/items/create_item`,
					{
						method: "POST",
						body: JSON.stringify(dataMes),
						headers: {
							"Content-Type": "application/json",
							Authorization: `Bearer ${token}`,
						},
						credentials: "include",
					},
				);
				if (!res.ok) {
					console.log("Сообщение отправить не получилось");
				}
			};
			if (String(id_user) === m.user_id) {
				console.log("я утуттутутуту");
				saveMessage();
			}

			if (refMes.current) {
				refMes.current.scrollTop = refMes.current.scrollHeight;
			}

			setMessages([...messages, m]);
			return;
		};

		conn.onclose = () => {
			console.log("ЧАТ ЗАКРЫЛСЯ");
		};
		conn.onerror = () => {
			console.log("АЙАЙ ОШИБОЧКА");
		};
		conn.onopen = () => {
			const getMessages = async () => {
				const res = await fetch(
					`http://localhost:8000/api/chats/${roomId[0]}/items/get_all_items`,
					{
						method: "GET",
						headers: {
							"Content-Type": "application/json",
							Authorization: `Bearer ${token}`,
						},
						credentials: "include",
					},
				);
				if (!res.ok) {
					console.log("Сообщения не удалось получить");
					return;
				}
				if (res.ok) {
					const chatItems = await res.json();
					console.log(chatItems, "ssssssssssss");
					const mess: Array<Message> = [];
					if (chatItems.data) {
						chatItems.data.forEach(
							(chatItem: {
								id: number;
								chatlist_id: string;
								username: string;
								description: string;
								user_id: number;
							}) => {
								const mes: Message = {
									id: chatItem.id,
									description: chatItem.description,
									user_id: String(chatItem.user_id),
									username: chatItem.username,
									chatlist_id: chatItem.chatlist_id,
								};
								mess.push(mes);
							},
						);
						console.log(mess);
						setMessages([...messages].concat(mess));
					}
				}
				if (refMes.current) {
					refMes.current.scrollTop = refMes.current.scrollHeight;
				}
			};
			getMessages();
		};
	}, [messages, conn, users]);

	const sendMessage = () => {
		if (message === "") return;
		if (conn === undefined || conn === null) {
			setActive(false);
			return;
		}

		conn.send(message);
		setMessage("");
	};

	if (active) {
		return (
			<div className={styles.chatApp}>
				<div className={styles.wrapper__chatApp__title}>
					<h1 className={styles.chatApp__title}></h1>
				</div>
				<div className={styles.chatApp__chatWindow}>
					<div className={styles.chatApp__messages} ref={refMes}>
						<div className={styles.messages}>
							<MessageBody messages={messages} />
						</div>
					</div>
					<div className={styles.chatApp__inputField}>
						<div className={styles.inputField__textarea}>
							<textarea
								ref={textarea}
								placeholder="Напишите что-нибудь..."
								value={message}
								onChange={(event) => {
									setMessage(event.target.value);
								}}
								className={styles.textarea__block}
							/>
						</div>
						<div className={styles.inputField__send}>
							<Button
								onClick={sendMessage}
								id={"sendMessage"}
								type={ButtonType.Icon}
								icon={<SendMessageIcon />}
							/>
						</div>
					</div>
				</div>
			</div>
		);
	} else {
		return <></>;
	}
};
