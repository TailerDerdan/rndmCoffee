import React from "react";
import { ComputerGamesActivity, TrendsActivity } from "../icons/icons";
import styles from "./hobbiesComponent.module.css";
import { AllActivities } from "../../data/data";
import { Headers } from "../../data/types";

type HobbiesType = {
	hobbies: Array<string> | null;
};

function getIconHobby(hobby: string) {
	let typeIcon: Headers = Headers.Trends;
	let isFind = false;
	for (const group of AllActivities.allActivities) {
		for (const hobbyFromGroup of group.activities) {
			if (hobbyFromGroup.description === hobby) {
				typeIcon = group.header;
				isFind = true;
				break;
			}
		}
		if (isFind) {
			break;
		}
	}
	switch (typeIcon) {
		case Headers.Trends:
			return <TrendsActivity />;
		case Headers.Games:
			return <ComputerGamesActivity />;
	}
}

export const HobbiesComponent = (props: HobbiesType) => {
	const { hobbies } = props;
	if (hobbies !== null) {
		hobbies.map((hobby, key) => {
			return (
				<div className={styles.wrapper__hobby} key={key}>
					{getIconHobby(hobby)}
				</div>
			);
		});
	} else {
		return <h2 className={styles.not__hobby}>Нет общих интересов</h2>;
	}
	return <></>;
};
