# Rocket_Elevators_Controllers

Contains the algorithm files for the elevator controllers for the New Rocket Elevator Solutions for both Residential and Commercial offers

// Residentiel 

    Au commencement l'utilisateur fait une requête pour demander l'ascenseur. Peut importe l'étage ou il se trouve, la premiere information que nous allons obtenir est sa position ( étage ) ainsi que sa direction ( UP ou DOWN ). Ceci correspond donc a la sequence "elevatorRequest".

    La deuxieme action qui sera donc nécessaire, après l'interaction de l'utilisateur avec le controleur de colonne a l'extérieur de l'ascenseur, sera donc de trouver l'ascenseur le plus adéquat. Ceci sera réalisé avec la séquence "findElevator". En prenant en compte certaines proprietés pour chacun des ascenceur du batiment (l'étage actuel, la direction de son mouvement et la difference d'étages entre celui-ci et l'utilisateur), je vais donc attribuer un système de pointage aux differents ascenceurs suivant certaines conditions.Par la suite, l'ascenceur obtenant le meilleur pointage et ayant la plus petite difference d'étages en l'utilisateur et lui-mème se verra séléctionner comme "selectedElevator". 

    Ensuite en se servant du "selectedElevator" établi a la sequence "findElevator", et de l'étage auquel l'utilisateur se situe, la sequence "moveElevator" va debuter. Le mouvement sera tout simplement défini en spécifiant que tant que l'étage actuel de notre ascenseur n'est pas égal a celui de l'utilisateur, l'ascenceur fera le mouvement approprié a la situation. Donc si l'étage de l'ascenceur est plus petit que celui de l'utilisateur, monte. Et ainsi de suite. Par la suite les portes vont tout simplement s'ouvrir un coup arriver a destination ou un coup que l'ascenceur ne sera plus en mouvement.

    Ensuite il suffit de recommencer le processus pour l'interaction avec l'utilisateur a l'interieur de l'ascenceur. En excluant bien sur la sequence "findElevator". 


// Commercial 

    Le batiment commercial va fonctionner d'une manière très semblable. A l'exception des colonnes. Étant donné le fait que nous avons plusieurs colonnes controlant plusieurs ascenceurs, nous devons ajouter une sequence avant "findElevator". 

    Après avoir établi un calcul pour spécifier combien d'étages sont déservit a chaque colonnes. Nous allons tout simplement nous servir de l'étages demandé au rez de chaussé par l'utilisateur. Ici le batiment Moderne est choisi comme modele et l'utilisateur doit spécifier au premier étage ou il veut aller. Le chiffre entré sera donc tout simplement analysé pour savoir a quelle liste d'étages deservit pour chaque colonnes il appartient. 

    Par exemple si l'étage spécifié par l'utilisateur est 8, l'algorithme va analyser chacune des colonnes et vérifier si #8 appartient a sa liste d'étages déservit. La sequence "findColumn" va ensuite retourner une colonne sélectionnée, "selectedColumn".

    Par la suite la séquence "findElevator" va determiné a l'aide de "selectedColumn" quel ascenceur envoyé selon certains parametres comme avec le controleur résidentiel et "bouger" l'ascenseur sélectionné. 