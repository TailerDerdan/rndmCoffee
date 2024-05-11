import React, { useEffect, useState } from "react";
import styles from "./homePage.module.css";
import { ManInChairIcon } from "../../../components/icons/icons";
import { MeetingForm } from "../../../components/formForMeet/meetingForm";
import { MeetCard } from "../../../components/meetCard/meetCard";
import { useLocalStorage } from "../../../hooks/useLocalStorage";
import { Chat } from "../../../components/chat/chat";

type Chat = {
	id: number;
	title: string;
};

export const HomePage = () => {
	const [activeForm, setActiveForm] = useState(false);
	const [activeChat, setActiveChat] = useState(false);
	const [chats, setChats] = useState<Array<Chat>>([]);
	const [conn, setConn] = useState<WebSocket | null>(null);

	let classesForIcon = styles.mainContent__iconMan;
	let classesForForm =
		styles.mainContent__form + " " + styles.wrapper__hidden;
	let classesForChat =
		styles.mainContent__chat + " " + styles.wrapper__hidden;
	if (activeForm) {
		classesForIcon += " " + styles.wrapper__hidden;
		classesForForm = styles.mainContent__form;
		classesForChat += " " + styles.wrapper__hidden;
	} else {
		classesForIcon = styles.mainContent__iconMan;
		classesForForm += " " + styles.wrapper__hidden;
	}
	if (activeChat) {
		classesForIcon += " " + styles.wrapper__hidden;
		classesForChat = styles.mainContent__chat;
	}

	const [token] = useLocalStorage({
		initialValue: "",
		key: "token",
	});

	const [profile_id, setProfileId] = useLocalStorage({
		initialValue: -1,
		key: "profile_id",
	});

	const [username, setUsername] = useLocalStorage({
		initialValue: "",
		key: "username",
	});

	const createHandler = async (chatId: number, title: string) => {
		const data = {
			Id: String(chatId),
			Name: title,
		};

		try {
			const res = await fetch(`http://localhost:8000/ws/createRoom`, {
				method: "POST",
				body: JSON.stringify(data),
				headers: {
					"Content-Type": "application/json",
				},
				credentials: "include",
			});
		} catch (err) {
			console.log(err);
		}
	};

	const getChats = async () => {
		try {
			const resGetChats = await fetch(
				"http://localhost:8000/api/chats/get_all_chats",
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${token}`,
					},
					credentials: "include",
				},
			);
			const resGetRooms = await fetch(
				"http://localhost:8000/ws/getRooms",
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${token}`,
					},
					credentials: "include",
				},
			);
			const dataFromDB = await resGetChats.json();
			const dataFromWS = await resGetRooms.json();
			if (resGetChats.ok && resGetRooms.ok) {
				const arrChats: Array<Chat> = [];
				dataFromDB.data.forEach(
					(chat: { id: string; clients: null; title: string }) => {
						arrChats.push({
							id: Number(chat.id),
							title: String(chat.title),
						});
					},
				);
				const rooms = dataFromWS.rooms;
				const roomsNotCreated = arrChats.filter((chat) => {
					const isCreated = rooms.find(
						(room: { id: string; name: string }) => {
							const idRoom = Number(room.id);
							const IdChat = chat.id;
							return idRoom === IdChat;
						},
					);
					if (isCreated === undefined) {
						return chat;
					}
				});
				roomsNotCreated.forEach((room) => {
					createHandler(room.id, room.title);
				});
				setChats(arrChats);
			}
		} catch (err) {
			console.log(err);
		}
	};

	useEffect(() => {
		getChats();
	}, []);

	const [id_user] = useLocalStorage({
		initialValue: "",
		key: "id_user",
	});

	const joinChat = (chatId: string) => {
		const ws = new WebSocket(
			`ws://127.0.0.1:8000/ws/joinRoom/${chatId}?userId=${id_user}&username=${username}`,
		);

		if (ws.OPEN) {
			setConn(ws);
			setActiveChat(true);
			return;
		}
	};

	return (
		<div className={styles.homePage}>
			<div className={styles.homePage__header}>
				<div className={styles.header__text}>
					<h1 className={styles.text}>Ваши встречи</h1>
				</div>
				<div className={styles.header__buttonOfMeetings}>
					<button
						className={styles.buttonOfMeetings}
						onClick={() => {
							setActiveForm(!activeForm);
						}}
					>
						Назначить встречу
					</button>
				</div>
			</div>
			<div className={styles.homePage__mainContent}>
				<div className={classesForIcon}>
					<ManInChairIcon />
				</div>
				<div className={styles.mainContent__meetings}>
					{chats.map((chat, index) => (
						<div className={styles.wrapper__meetCard} key={index}>
							<MeetCard
								chatId={chat.id}
								key={index}
								joinChat={joinChat}
							/>
						</div>
					))}
				</div>
				<div className={classesForForm}>
					<MeetingForm
						active={activeForm}
						setActive={setActiveForm}
						getChats={getChats}
						createHandler={createHandler}
					/>
				</div>
				<div className={classesForChat}>
					<Chat
						setActive={setActiveChat}
						active={activeChat}
						conn={conn}
					/>
				</div>
			</div>
		</div>
	);
};