local CharacterController = {
    model = Instance,
    humanoid = Instance,
    walkSpeed = number,
    jumpPower = number,
    moveDirection = Vector3,
    isJumping = boolean,
    states = table
}

local CharacterAnimation = {
    animator = Instance,
    animations = table,
    currentAnim = string
}

local CharacterInput = {
    movement = Vector2,
    jump = boolean,
    cameraOffset = CFrame
}

world:system({
    name = "CharacterMovement",
    parameters = {deltaTime = 0},
    query = {
        all = {
            CharacterController,
            CharacterInput
        },
    },
    frequency = <skipped>,
    priority = 100,
    callback = function(entity, components, deltaTime)
        for _ , entity in ipairs ( entities ) do local controller = entity . CharacterController local input = entity . CharacterInput local camera = workspace . CurrentCamera local lookVector = camera . CFrame . LookVector local rightVector = camera . CFrame . RightVector local moveVector = Vector3 . new ( input . movement . X * rightVector . X + input . movement . Y * lookVector . X , 0 , input . movement . X * rightVector . Z + input . movement . Y * lookVector . Z ) . Unit * controller . walkSpeed if controller . humanoid then controller . humanoid : Move ( moveVector ) end if input . jump and not controller . isJumping then controller . isJumping = true if controller . humanoid then controller . humanoid . Jump = true end end end
    end
})

world:system({
    name = "CharacterAnimator",
    query = {
        all = {
            CharacterController,
            CharacterAnimation
        },
    },
    callback = function(entity, components)
        for _ , entity in ipairs ( entities ) do local controller = entity . CharacterController local anim = entity . CharacterAnimation if not anim . animator then continue end local newAnim = idle if controller . isJumping then newAnim = jump elseif controller . moveDirection . Magnitude > 0.1 then newAnim = walk end if newAnim ~ = anim . currentAnim then anim . currentAnim = newAnim anim . animator : LoadAnimation ( anim . animations [ newAnim ] ) : Play ( ) end end
    end
})

world:system({
    name = "CameraController",
    parameters = {deltaTime = 0, smoothing = 0.2},
    query = {
        all = {
            CharacterController,
            CharacterInput
        },
    },
    callback = function(entity, components, deltaTime, smoothing)
        for _ , entity in ipairs ( entities ) do local controller = entity . CharacterController local input = entity . CharacterInput if controller . model then local camera = workspace . CurrentCamera local targetCFrame = controller . model : GetPivot ( ) * input . cameraOffset camera . CFrame = camera . CFrame : Lerp ( targetCFrame , smoothing ) end end
    end
})