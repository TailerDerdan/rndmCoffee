import React from "react";
import { ProfileIconMeetCard } from "../icons/icons";
import styles from "./photo.module.css";

type PhotosType = {
	photos: Array<string>;
};

export const PhotosComponent = (props: PhotosType) => {
	const { photos } = props;
	return (
		<>
			{photos.map((photo, key) => {
				if (photo === "none") {
					return (
						<div className={styles.wrapper__avatar} key={key}>
							<ProfileIconMeetCard />
						</div>
					);
				} else {
					return (
						<div className={styles.wrapper__avatar} key={key}>
							<img src="photo" alt="avatar" />
						</div>
					);
				}
			})}
		</>
	);
};
