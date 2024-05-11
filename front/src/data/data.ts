import { GroupActivity, Headers, TypeAllActivities } from "./types";

export const Trends: GroupActivity = {
	header: Headers.Trends,
	activities: [
		{ description: "NFL" },
		{ description: "NBA" },
		{ description: "Мировые новости" },
		{ description: "ChatGPT" },
		{ description: "One Piece" },
		{ description: "Midjourney" },
		{ description: "Cплетни о знаменитостях" },
	],
};

export const Games: GroupActivity = {
	header: Headers.Games,
	activities: [
		{ description: "Call of Duty" },
		{ description: "Baldur’s Gate 3" },
		{ description: "Minecraft" },
		{ description: "Playstation" },
		{ description: "Genshin Impact" },
		{ description: "GTA" },
		{ description: "Sims" },
		{ description: "Terraria" },
		{ description: "Red Dead Redemption" },
	],
};

export const Art: GroupActivity = {
	header: Headers.Art,
	activities: [
		{ description: "Фотография" },
		{ description: "Видеосъемка" },
		{ description: "Дизайн" },
		{ description: "Макияж" },
		{ description: "Рукоделие" },
		{ description: "Танцы" },
		{ description: "Пения" },
		{ description: "Музыка" },
		{ description: "Ведение блога" },
		{ description: "Рисование" },
	],
};

// export const ActiveLifeStyle: GroupActivity = {
// 	header: Headers.ActiveLifestyle,
// 	activities: [
// 		{ description: "Космос" },
// 		{ description: "Природа" },
// 		{ description: "Новости про науку" },
// 		{ description: "Крутые гайды" },
// 		{ description: "История" },
// 		{ description: "География" },
// 		{ description: "Философия" },
// 		{ description: "Математика" },
// 		{ description: "Computer Science" },
// 		{ description: "Экономика" },
// 	],
// };

export const AllActivities: TypeAllActivities = {
	allActivities: [Trends, Games],
};
