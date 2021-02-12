# Oh Mummy GO
Pequeño remake del juego [Oh Mummy](https://www.youtube.com/watch?v=Ls5AGwkRNz0) realizado en golang y con la librería gráfica [ebiten](https://ebiten.org/).

## Plataforma de desarrollo.
* Ubuntu 20.04
* golang 1.15.8
* ebiten v1.12.7

# Progreso.
## TODO
* Finalizar el juego:
    * cargando otro nivel (positvo).

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