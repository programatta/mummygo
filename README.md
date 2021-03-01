# Oh Mummy GO
Pequeño remake del juego [Oh Mummy](https://www.youtube.com/watch?v=Ls5AGwkRNz0) realizado en golang y con la librería gráfica [ebiten](https://ebiten.org/).

## Plataforma de desarrollo.
* Ubuntu 20.04
* golang 1.15.8
* ebiten v1.12.7

# Progreso.
## TODO
* FX hechizo saliendo
* Añadir estados en el propio juego que nos permita pasar de playing, gameover, nextlevel

## 20210301.
* Se añade datos sobre los niveles de juego
* Se añade funcionalidad para la carga de niveles y preparar el juego para un nuevo nivel.
* Se añade funcionalidad para el estado de game over y volver a jugar desde el inicio al pulsar jugar desde el menu ppal.
* Se añade funcionalidad para detectar el último nivel y pasar a un estado de victoria y vuelta al menu principal.
* Se añade sonido de estado de victoria.
* Bug fixed: Al estar hechizado las momias se paran detrás del judagor.
## 20210218.
* Error en audio (necesita installar alsa)
~~~
$sudo apt install libasound2-dev
~~~

* Se añade funcionalidad para un cargador de sonidos (de momento wav).
* Se añaden sonidos:
    * de fondo en el gameplay
    * apertura de puerta en la tumba
    * apertura de puerta principal
    * pasos del jugador
    * muerte del judagor
    * momia saliendo
    * momia gruñendo
    * momia muriendo
    * hechizo golpeando al jugador
    * hechizo sin efecto al golpear al jugador
    * coger poción
    * coger llave y papiro
    * pocion saliendo
    * items saliendo

## 20210212.
* Se añade funcionalidad para un cargador de fuentes.
* Pequeña refactorización.
* Se añade funconalidad para los estado de Menu Principal y Créditos
* Pequeña funcionalidad para conteo de puntos.

## 20210211.
* Se añade funcionalidad para el movimiento de las mominas y sigan al jugador usando un algoritmo A* (pathfinding).
* Se añade funcionalidad para nuevo enemigo (hechizo) que también hace uso de pathfinding.

## 20210210.
* Estados internos para el gameplay y player
    * carga un nuevo nivel y preparar el escenario (wip).
    * abandonar el nivel de forma bonita por parte del player (Se queda la pantalla en negro).
* Se añade funcionalidad para futuros estados de juego
    * Estado actual: GamePlay
    * Futuros estados: MainMenu y Credits.
* Se añade funcionalidad para cuando el player sea alcanzado por una momia y aun tenga vidas sea situado en el punto inicial y durante unos segundo no puede ser comido de nuevo.

## 20210209.
* Se añade funcionalidad para apertura de tumbas y que se muestren los objetos que hay dentro.
* Se añade funcionalidad para que la aparición de objetos sea mas bonita a través de estados (muy básico).
* Se añade funcionalidad para cargar fuentes ttf y mostrarlas en el UI
* Se recolectan objetos y funcionalidad para el que están destinados (y actualización del UI).
* Se añade funcionalidad básica de Game Over (Se queda la pantalla en negro).
* Se añade funcionalidad básica de nivel completado (Se abre la puerta principal).

## 20210208.
Preparación del proyecto y se añade funcionalidad básica para:
* Cargar fichero de assets (spriteSheet).
* Realizar el pintado del panel de juego.
* Capturar eventos de teclado.
* Añadir pequeña funcionalidad en la clase jugador y momia.


# Recursos.
## Fonts.
* [Barcade-brawl]https://www.fontspace.com/barcade-brawl-font-f31534) de [Pixel Kitchen](https://www.fontspace.com/pixel-kitchen)
* [Karmatic-arcade](https://www.1001freefonts.com/karmatic-arcade.font) de [Vic Fieger](https://www.1001freefonts.com/designer-vic-fieger-fontlisting.php)

## Código.
* Algoritmo de pathfinding de [Alex Plugaru](ttps://github.com/xarg/gopathfinding)

## FX/SOunds
* echar un ojo a: 
    * https://www.zapsplat.com/sound-effect-category/monsters-and-creatures/page/3/
    * Asset store de unity.

* sonido puerta abriendo: https://freesound.org/people/gabisaraceni/sounds/96964/

* sonido puerta metal abriendo: https://freesound.org/people/nebyoolae/sounds/267692/

* sonido de pasos: https://freesound.org/people/Mativve/sounds/414335/

* sonido muerte jugador: https://freesound.org/people/Haramir/sounds/404014/

* sonido hechizo alcanza jugador: https://freesound.org/people/EminYILDIRIM/sounds/550267/
- https://freesound.org/people/LittleRobotSoundFactory/sounds/270409/

* sonido hechizo sin efecto al alcanzar jugador: https://freesound.org/people/spookymodem/sounds/249819/

* sonido cuando sale la momia: https://freesound.org/people/d761747/sounds/503096/
- https://freesound.org/people/Speedenza/sounds/168180/

* sonido momia rugiendo: https://freesound.org/people/Antimsounds/sounds/556048/
- https://freesound.org/people/Slaking_97/sounds/332023/

* sonido momia muriendo: https://freesound.org/people/Michel88/sounds/76957/

* sonido pocion saliendo: https://freesound.org/people/bongmoth/sounds/156566/

* sonido coger la pocion: https://freesound.org/people/Jamius/sounds/41529/

* sonido item saliendo: https://freesound.org/people/Mrthenoronha/sounds/388895/

* sonido coger item: https://freesound.org/people/MakoFox/sounds/126412/

* sonido ambiente main menu y creditos: https://freesound.org/people/Jedo/sounds/397009/

* sonido ambiente juego : https://freesound.org/people/TheFrohman/sounds/206043/

* sonido de game over : https://freesound.org/people/MATRIXXX_/sounds/365782/

* sonido de nivel finalizado: https://freesound.org/people/jalastram/sounds/317593/

* sonido de juego acabado: https://freesound.org/people/LittleRobotSoundFactory/sounds/270402/
  Está a 48000Hz -> 44100Hz (https://audio.online-convert.com/convert-to-wav)