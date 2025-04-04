ALTER TABLE "MapElements" DROP CONSTRAINT "MapElements_elementId_fkey";
ALTER TABLE "MapElements" DROP CONSTRAINT "MapElements_mapId_fkey";
ALTER TABLE "spaceElements" DROP CONSTRAINT "spaceElements_elementId_fkey";
ALTER TABLE "spaceElements" DROP CONSTRAINT "spaceElements_spaceId_fkey";
ALTER TABLE "Space" DROP CONSTRAINT "Space_creatorId_fkey";
ALTER TABLE "User" DROP CONSTRAINT "User_avatarId_fkey";

DROP INDEX "Avatar_id_key";
DROP INDEX "MapElements_id_key";
DROP INDEX "Map_id_key";
DROP INDEX "Element_id_key";
DROP INDEX "spaceElements_id_key";
DROP INDEX "Space_id_key";
DROP INDEX "User_username_key";
DROP INDEX "User_id_key";

DROP TABLE "Avatar";
DROP TABLE "MapElements";
DROP TABLE "Map";
DROP TABLE "Element";
DROP TABLE "spaceElements";
DROP TABLE "Space";
DROP TABLE "User";

DROP TYPE "Role";