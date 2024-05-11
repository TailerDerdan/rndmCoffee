import React, { useContext, useEffect, useMemo, useState } from "react";
import styles from "./meetCard.module.css";
import { LineSeparatorMeetCard } from "../icons/icons";
import { useLocalStorage } from "../../hooks/useLocalStorage";
import { PhotosComponent } from "./photos";
import { HobbiesComponent } from "./hobbiesComponent";

type MeetCardProps = {
	chatId: number;
	joinChat: (chatId: string) => void;
};

export const MeetCard = (props: MeetCardProps) => {
	const { chatId, joinChat } = props;
	const [photos, setPhotos] = useState<Array<string>>([]);
	const [hobbies, setHobbies] = useState<Array<string> | null>([]);

	const [token] = useLocalStorage({
		initialValue: "",
		key: "token",
	});

	useEffect(() => {
		async function GetChat() {
			const res = await fetch(
				`http://localhost:8000/api/chats/get_info_for_chat/${chatId}`,
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${token}`,
					},
					credentials: "include",
				},
			);
			if (res.ok) {
				const chat = await res.json();
				const photos: Array<string> = [];
				for (const key in chat.users_info) {
					photos.push(chat.users_info[key]);
				}
				setHobbies(chat.data);
				setPhotos(photos);
			}
		}
		GetChat();
	}, []);

	return (
		<div
			className={styles.wrapper__meetCard}
			onClick={(event) => {
				event.preventDefault();
				setTimeout(() => {
					joinChat(String(chatId));
				}, 250);
			}}
		>
			<div className={styles.wrapper__users}>
				<PhotosComponent photos={photos} />
			</div>
			<div className={styles.wrapper__text}>
				<h1 className={styles.text}>Новая встреча</h1>
			</div>
			<div className={styles.wrapper__activities}>
				<HobbiesComponent hobbies={hobbies} />
			</div>
		</div>
	);
};
