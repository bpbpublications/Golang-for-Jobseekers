USE `application`;

CREATE TABLE IF NOT EXISTS `revenues` (
  `datetime` varchar(100) NOT NULL,      
  `revenue` float NOT NULL default 0,
   PRIMARY KEY  (`datetime`)
);
