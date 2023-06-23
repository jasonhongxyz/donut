# donut

Credits: https://www.a1k0n.net/2011/07/20/donut-math.html

![donut](./donut.gif)

## Notes
1. Create a circle offset from $(0,0)$ by $R_2$

$$
(x, y, z) 
= (R_2, 0, 0) + (R_1 cos \theta, R_1 sin \theta, 0) 
= (R_2 + R_1 cos \theta, R_1 sin \theta, 0)
$$

2. A donut (torus) is a solid of revolution. Multiply the above circle by the $R_y(\phi)$ [rotation matrix](https://en.wikipedia.org/wiki/Rotation_matrix) by another angle $\phi$.

```math
$$ (R_2 + R_1 cos \theta, R_1 sin \theta, 0) * \begin{bmatrix}cos \theta & 0 & sin \theta \\ 0 & 1 & 0 \\ -sin \theta & 0 & cos \theta \end{bmatrix} $$
```

$$ = ((R_2 +R_1cos\theta)cos\theta, R_1sin\theta, (R_2 + R_1cos\theta)sin\theta) $$

3. Multiply by the other rotation matrices $R_x(A)$ and $R_z(B)$ to rotate in the $x$ and $z$ dimensions.

```math
$$ (R_2 + R_1 cos \theta, R_1 sin \theta, 0) * \begin{bmatrix}cos \theta & 0 & sin \theta \\ 0 & 1 & 0 \\ -sin \theta & 0 & cos \theta \end{bmatrix} * \begin{bmatrix} 1 & 0 & 0 \\ 0 & cos A & -sin A \\ 0 & sin A & cos A \end{bmatrix} * \begin{bmatrix} cos B & -sin B & 0 \\ sin B & cos B & 0 \\ 0 & 0 & 1\end{bmatrix}$$
```

4. To map 3D space to 2D, we can find proportions for $(x, y, z)$ to $(x', y')$, where...

$$ \frac{x'}{z'} = \frac{x}{z} $$
$$ x' = \frac{xz'}{z}$$

$$ \frac{y'}{z'} = \frac{y}{z} $$
$$ y' = \frac{yz'}{z}$$

5. Move the donut further away from viewer by adjusting denominator $z$ by another constant $K_2$.

$$ (x', y') = (\frac{xz'}{K_2+z}, \frac{yz'}{K_2 + z}) $$

6. For luminance, we first need the [surface normal](https://en.wikipedia.org/wiki/Normal_(geometry)) of the donut. Similar to above, except this time, rotate on $(cos \theta, sin \theta, 0)$, or the center of the torus.
```math
$$ (N_x, N_y, N_z) = (cos \theta, sin \theta, 0) * \begin{bmatrix}cos \theta & 0 & sin \theta \\ 0 & 1 & 0 \\ -sin \theta & 0 & cos \theta \end{bmatrix} * \begin{bmatrix} 1 & 0 & 0 \\ 0 & cos A & -sin A \\ 0 & sin A & cos A \end{bmatrix} * \begin{bmatrix} cos B & -sin B & 0 \\ sin B & cos B & 0 \\ 0 & 0 & 1\end{bmatrix}$$
```

7. If we choose location $(0, 1, -1)$ as the light source point, we multiply the above matrix by $(0, 1, -1)$.

8. Finally, pick values for $R_1, R_2, K_1, K_2$.
